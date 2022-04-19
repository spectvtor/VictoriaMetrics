// Code generated by qtc from "query_range_response.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line app/vmselect/prometheus/query_range_response.qtpl:1
package prometheus

//line app/vmselect/prometheus/query_range_response.qtpl:1
import (
	"github.com/VictoriaMetrics/VictoriaMetrics/app/vmselect/netstorage"
)

// QueryRangeResponse generates response for /api/v1/query_range.See https://prometheus.io/docs/prometheus/latest/querying/api/#range-queries

//line app/vmselect/prometheus/query_range_response.qtpl:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line app/vmselect/prometheus/query_range_response.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line app/vmselect/prometheus/query_range_response.qtpl:8
func StreamQueryRangeResponse(qw422016 *qt422016.Writer, rs []netstorage.Result, trace string) {
//line app/vmselect/prometheus/query_range_response.qtpl:8
	qw422016.N().S(`{"status":"success",`)
//line app/vmselect/prometheus/query_range_response.qtpl:11
	if trace != "" {
//line app/vmselect/prometheus/query_range_response.qtpl:11
		qw422016.N().S(`"trace":`)
//line app/vmselect/prometheus/query_range_response.qtpl:12
		qw422016.N().S(trace)
//line app/vmselect/prometheus/query_range_response.qtpl:12
		qw422016.N().S(`,`)
//line app/vmselect/prometheus/query_range_response.qtpl:13
	}
//line app/vmselect/prometheus/query_range_response.qtpl:13
	qw422016.N().S(`"data":{"resultType":"matrix","result":[`)
//line app/vmselect/prometheus/query_range_response.qtpl:17
	if len(rs) > 0 {
//line app/vmselect/prometheus/query_range_response.qtpl:18
		streamqueryRangeLine(qw422016, &rs[0])
//line app/vmselect/prometheus/query_range_response.qtpl:19
		rs = rs[1:]

//line app/vmselect/prometheus/query_range_response.qtpl:20
		for i := range rs {
//line app/vmselect/prometheus/query_range_response.qtpl:20
			qw422016.N().S(`,`)
//line app/vmselect/prometheus/query_range_response.qtpl:21
			streamqueryRangeLine(qw422016, &rs[i])
//line app/vmselect/prometheus/query_range_response.qtpl:22
		}
//line app/vmselect/prometheus/query_range_response.qtpl:23
	}
//line app/vmselect/prometheus/query_range_response.qtpl:23
	qw422016.N().S(`]}}`)
//line app/vmselect/prometheus/query_range_response.qtpl:27
}

//line app/vmselect/prometheus/query_range_response.qtpl:27
func WriteQueryRangeResponse(qq422016 qtio422016.Writer, rs []netstorage.Result, trace string) {
//line app/vmselect/prometheus/query_range_response.qtpl:27
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/prometheus/query_range_response.qtpl:27
	StreamQueryRangeResponse(qw422016, rs, trace)
//line app/vmselect/prometheus/query_range_response.qtpl:27
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/prometheus/query_range_response.qtpl:27
}

//line app/vmselect/prometheus/query_range_response.qtpl:27
func QueryRangeResponse(rs []netstorage.Result, trace string) string {
//line app/vmselect/prometheus/query_range_response.qtpl:27
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/prometheus/query_range_response.qtpl:27
	WriteQueryRangeResponse(qb422016, rs, trace)
//line app/vmselect/prometheus/query_range_response.qtpl:27
	qs422016 := string(qb422016.B)
//line app/vmselect/prometheus/query_range_response.qtpl:27
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/prometheus/query_range_response.qtpl:27
	return qs422016
//line app/vmselect/prometheus/query_range_response.qtpl:27
}

//line app/vmselect/prometheus/query_range_response.qtpl:29
func streamqueryRangeLine(qw422016 *qt422016.Writer, r *netstorage.Result) {
//line app/vmselect/prometheus/query_range_response.qtpl:29
	qw422016.N().S(`{"metric":`)
//line app/vmselect/prometheus/query_range_response.qtpl:31
	streammetricNameObject(qw422016, &r.MetricName)
//line app/vmselect/prometheus/query_range_response.qtpl:31
	qw422016.N().S(`,"values":`)
//line app/vmselect/prometheus/query_range_response.qtpl:32
	streamvaluesWithTimestamps(qw422016, r.Values, r.Timestamps)
//line app/vmselect/prometheus/query_range_response.qtpl:32
	qw422016.N().S(`}`)
//line app/vmselect/prometheus/query_range_response.qtpl:34
}

//line app/vmselect/prometheus/query_range_response.qtpl:34
func writequeryRangeLine(qq422016 qtio422016.Writer, r *netstorage.Result) {
//line app/vmselect/prometheus/query_range_response.qtpl:34
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/prometheus/query_range_response.qtpl:34
	streamqueryRangeLine(qw422016, r)
//line app/vmselect/prometheus/query_range_response.qtpl:34
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/prometheus/query_range_response.qtpl:34
}

//line app/vmselect/prometheus/query_range_response.qtpl:34
func queryRangeLine(r *netstorage.Result) string {
//line app/vmselect/prometheus/query_range_response.qtpl:34
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/prometheus/query_range_response.qtpl:34
	writequeryRangeLine(qb422016, r)
//line app/vmselect/prometheus/query_range_response.qtpl:34
	qs422016 := string(qb422016.B)
//line app/vmselect/prometheus/query_range_response.qtpl:34
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/prometheus/query_range_response.qtpl:34
	return qs422016
//line app/vmselect/prometheus/query_range_response.qtpl:34
}
