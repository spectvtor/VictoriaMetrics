// Code generated by qtc from "json.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line lib/tracer/json.qtpl:2
package tracer

//line lib/tracer/json.qtpl:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line lib/tracer/json.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line lib/tracer/json.qtpl:2
func streamctxToJson(qw422016 *qt422016.Writer, ctx *Context) {
//line lib/tracer/json.qtpl:2
	qw422016.N().S(`{"id": "`)
//line lib/tracer/json.qtpl:4
	qw422016.N().S(ctx.ID())
//line lib/tracer/json.qtpl:4
	qw422016.N().S(`","msg":`)
//line lib/tracer/json.qtpl:5
	qw422016.N().Q(ctx.msg)
//line lib/tracer/json.qtpl:5
	qw422016.N().S(`,"elapsed": "`)
//line lib/tracer/json.qtpl:6
	qw422016.N().D(int(ctx.updated.Sub(ctx.start).Milliseconds()))
//line lib/tracer/json.qtpl:6
	qw422016.N().S(`ms","subtrace": [`)
//line lib/tracer/json.qtpl:8
	for i, sub := range ctx.subContexts {
//line lib/tracer/json.qtpl:9
		streamsubCtxToJson(qw422016, sub)
//line lib/tracer/json.qtpl:9
		if len(ctx.subContexts) != i+1 {
//line lib/tracer/json.qtpl:9
			qw422016.N().S(`,`)
//line lib/tracer/json.qtpl:9
		}
//line lib/tracer/json.qtpl:10
	}
//line lib/tracer/json.qtpl:10
	qw422016.N().S(`]}`)
//line lib/tracer/json.qtpl:13
}

//line lib/tracer/json.qtpl:13
func writectxToJson(qq422016 qtio422016.Writer, ctx *Context) {
//line lib/tracer/json.qtpl:13
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/tracer/json.qtpl:13
	streamctxToJson(qw422016, ctx)
//line lib/tracer/json.qtpl:13
	qt422016.ReleaseWriter(qw422016)
//line lib/tracer/json.qtpl:13
}

//line lib/tracer/json.qtpl:13
func ctxToJson(ctx *Context) string {
//line lib/tracer/json.qtpl:13
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/tracer/json.qtpl:13
	writectxToJson(qb422016, ctx)
//line lib/tracer/json.qtpl:13
	qs422016 := string(qb422016.B)
//line lib/tracer/json.qtpl:13
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/tracer/json.qtpl:13
	return qs422016
//line lib/tracer/json.qtpl:13
}

//line lib/tracer/json.qtpl:15
func streamsubCtxToJson(qw422016 *qt422016.Writer, ctx *Context) {
//line lib/tracer/json.qtpl:15
	qw422016.N().S(`{"msg":`)
//line lib/tracer/json.qtpl:17
	qw422016.N().Q(ctx.msg)
//line lib/tracer/json.qtpl:17
	qw422016.N().S(`,"elapsed": "`)
//line lib/tracer/json.qtpl:18
	qw422016.N().D(int(ctx.updated.Sub(ctx.start).Milliseconds()))
//line lib/tracer/json.qtpl:18
	qw422016.N().S(`ms","subtrace": [`)
//line lib/tracer/json.qtpl:20
	for i, sub := range ctx.subContexts {
//line lib/tracer/json.qtpl:21
		streamsubCtxToJson(qw422016, sub)
//line lib/tracer/json.qtpl:21
		if len(ctx.subContexts) != i+1 {
//line lib/tracer/json.qtpl:21
			qw422016.N().S(`,`)
//line lib/tracer/json.qtpl:21
		}
//line lib/tracer/json.qtpl:22
	}
//line lib/tracer/json.qtpl:22
	qw422016.N().S(`]}`)
//line lib/tracer/json.qtpl:25
}

//line lib/tracer/json.qtpl:25
func writesubCtxToJson(qq422016 qtio422016.Writer, ctx *Context) {
//line lib/tracer/json.qtpl:25
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/tracer/json.qtpl:25
	streamsubCtxToJson(qw422016, ctx)
//line lib/tracer/json.qtpl:25
	qt422016.ReleaseWriter(qw422016)
//line lib/tracer/json.qtpl:25
}

//line lib/tracer/json.qtpl:25
func subCtxToJson(ctx *Context) string {
//line lib/tracer/json.qtpl:25
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/tracer/json.qtpl:25
	writesubCtxToJson(qb422016, ctx)
//line lib/tracer/json.qtpl:25
	qs422016 := string(qb422016.B)
//line lib/tracer/json.qtpl:25
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/tracer/json.qtpl:25
	return qs422016
//line lib/tracer/json.qtpl:25
}