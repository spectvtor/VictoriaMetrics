package vmimport

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"sync"

	"github.com/VictoriaMetrics/VictoriaMetrics/app/vminsert/common"
	"github.com/VictoriaMetrics/VictoriaMetrics/app/vminsert/concurrencylimiter"
	"github.com/VictoriaMetrics/VictoriaMetrics/app/vminsert/netstorage"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/auth"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/bytesutil"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/storage"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/tenantmetrics"
	"github.com/VictoriaMetrics/metrics"
)

var maxLineLen = flag.Int("import.maxLineLen", 100*1024*1024, "The maximum length in bytes of a single line accepted by /api/v1/import")

var (
	rowsInserted  = tenantmetrics.NewCounterMap(`vm_rows_inserted_total{type="vmimport"}`)
	rowsPerInsert = metrics.NewSummary(`vm_rows_per_insert{type="vmimport"}`)
)

// InsertHandler processes `/api/v1/import` request.
//
// See https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6
func InsertHandler(at *auth.Token, req *http.Request) error {
	return concurrencylimiter.Do(func() error {
		return insertHandlerInternal(at, req)
	})
}

func insertHandlerInternal(at *auth.Token, req *http.Request) error {
	readCalls.Inc()

	r := req.Body
	if req.Header.Get("Content-Encoding") == "gzip" {
		zr, err := common.GetGzipReader(r)
		if err != nil {
			return fmt.Errorf("cannot read gzipped vmimport data: %s", err)
		}
		defer common.PutGzipReader(zr)
		r = zr
	}

	ctx := getPushCtx()
	defer putPushCtx(ctx)
	for ctx.Read(r) {
		if err := ctx.InsertRows(at); err != nil {
			return err
		}
	}
	return ctx.Error()
}

func (ctx *pushCtx) InsertRows(at *auth.Token) error {
	rows := ctx.Rows.Rows
	ic := &ctx.Common
	ic.Reset()
	rowsTotal := 0
	for i := range rows {
		r := &rows[i]
		ic.Labels = ic.Labels[:0]
		for j := range r.Tags {
			tag := &r.Tags[j]
			ic.AddLabelBytes(tag.Key, tag.Value)
		}
		ic.MetricNameBuf = storage.MarshalMetricNameRaw(ic.MetricNameBuf[:0], at.AccountID, at.ProjectID, ic.Labels)
		storageNodeIdx := ic.GetStorageNodeIdx(at, ic.Labels)
		values := r.Values
		timestamps := r.Timestamps
		_ = timestamps[len(values)-1]
		for j, value := range values {
			timestamp := timestamps[j]
			if err := ic.WriteDataPointExt(at, storageNodeIdx, ic.MetricNameBuf, timestamp, value); err != nil {
				return err
			}
		}
		rowsTotal += len(values)
	}
	rowsInserted.Get(at).Add(rowsTotal)
	rowsPerInsert.Update(float64(rowsTotal))
	return ic.FlushBufs()
}

func (ctx *pushCtx) Read(r io.Reader) bool {
	if ctx.err != nil {
		return false
	}
	ctx.reqBuf, ctx.tailBuf, ctx.err = common.ReadLinesBlockExt(r, ctx.reqBuf, ctx.tailBuf, *maxLineLen)
	if ctx.err != nil {
		if ctx.err != io.EOF {
			readErrors.Inc()
			ctx.err = fmt.Errorf("cannot read vmimport data: %s", ctx.err)
		}
		return false
	}
	ctx.Rows.Unmarshal(bytesutil.ToUnsafeString(ctx.reqBuf))
	return true
}

var (
	readCalls  = metrics.NewCounter(`vm_read_calls_total{name="vmimport"}`)
	readErrors = metrics.NewCounter(`vm_read_errors_total{name="vmimport"}`)
)

type pushCtx struct {
	Rows   Rows
	Common netstorage.InsertCtx

	reqBuf  []byte
	tailBuf []byte

	err error
}

func (ctx *pushCtx) Error() error {
	if ctx.err == io.EOF {
		return nil
	}
	return ctx.err
}

func (ctx *pushCtx) reset() {
	ctx.Rows.Reset()
	ctx.Common.Reset()
	ctx.reqBuf = ctx.reqBuf[:0]
	ctx.tailBuf = ctx.tailBuf[:0]

	ctx.err = nil
}

func getPushCtx() *pushCtx {
	select {
	case ctx := <-pushCtxPoolCh:
		return ctx
	default:
		if v := pushCtxPool.Get(); v != nil {
			return v.(*pushCtx)
		}
		return &pushCtx{}
	}
}

func putPushCtx(ctx *pushCtx) {
	ctx.reset()
	select {
	case pushCtxPoolCh <- ctx:
	default:
		pushCtxPool.Put(ctx)
	}
}

var pushCtxPool sync.Pool
var pushCtxPoolCh = make(chan *pushCtx, runtime.GOMAXPROCS(-1))
