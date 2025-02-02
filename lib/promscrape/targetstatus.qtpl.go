// Code generated by qtc from "targetstatus.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line lib/promscrape/targetstatus.qtpl:1
package promscrape

//line lib/promscrape/targetstatus.qtpl:1
import (
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/prompbmarshal"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/promrelabel"
	"net/url"
	"time"
)

//line lib/promscrape/targetstatus.qtpl:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line lib/promscrape/targetstatus.qtpl:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line lib/promscrape/targetstatus.qtpl:10
func StreamTargetsResponsePlain(qw422016 *qt422016.Writer, jts []jobTargetsStatuses, emptyJobs []string, showOriginLabels, showOnlyUnhealthy bool, err error) {
//line lib/promscrape/targetstatus.qtpl:12
	if err != nil {
//line lib/promscrape/targetstatus.qtpl:13
		qw422016.N().S(err.Error())
//line lib/promscrape/targetstatus.qtpl:14
		return
//line lib/promscrape/targetstatus.qtpl:15
	}
//line lib/promscrape/targetstatus.qtpl:17
	for _, js := range jts {
//line lib/promscrape/targetstatus.qtpl:18
		if showOnlyUnhealthy && js.upCount == js.targetsTotal {
//line lib/promscrape/targetstatus.qtpl:18
			continue
//line lib/promscrape/targetstatus.qtpl:18
		}
//line lib/promscrape/targetstatus.qtpl:18
		qw422016.N().S(`job=`)
//line lib/promscrape/targetstatus.qtpl:19
		qw422016.N().Q(js.job)
//line lib/promscrape/targetstatus.qtpl:19
		qw422016.N().S(`(`)
//line lib/promscrape/targetstatus.qtpl:19
		qw422016.N().D(js.upCount)
//line lib/promscrape/targetstatus.qtpl:19
		qw422016.N().S(`/`)
//line lib/promscrape/targetstatus.qtpl:19
		qw422016.N().D(js.targetsTotal)
//line lib/promscrape/targetstatus.qtpl:19
		qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:19
		qw422016.N().S(`up)`)
//line lib/promscrape/targetstatus.qtpl:20
		qw422016.N().S(`
`)
//line lib/promscrape/targetstatus.qtpl:21
		for _, ts := range js.targetsStatus {
//line lib/promscrape/targetstatus.qtpl:22
			if showOnlyUnhealthy && ts.up {
//line lib/promscrape/targetstatus.qtpl:22
				continue
//line lib/promscrape/targetstatus.qtpl:22
			}
//line lib/promscrape/targetstatus.qtpl:23
			qw422016.N().S("\t")
//line lib/promscrape/targetstatus.qtpl:23
			qw422016.N().S(`state=`)
//line lib/promscrape/targetstatus.qtpl:24
			if ts.up {
//line lib/promscrape/targetstatus.qtpl:24
				qw422016.N().S(`up`)
//line lib/promscrape/targetstatus.qtpl:24
			} else {
//line lib/promscrape/targetstatus.qtpl:24
				qw422016.N().S(`down`)
//line lib/promscrape/targetstatus.qtpl:24
			}
//line lib/promscrape/targetstatus.qtpl:24
			qw422016.N().S(`,`)
//line lib/promscrape/targetstatus.qtpl:24
			qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:24
			qw422016.N().S(`endpoint=`)
//line lib/promscrape/targetstatus.qtpl:25
			qw422016.N().S(ts.sw.Config.ScrapeURL)
//line lib/promscrape/targetstatus.qtpl:25
			qw422016.N().S(`,`)
//line lib/promscrape/targetstatus.qtpl:25
			qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:25
			qw422016.N().S(`labels=`)
//line lib/promscrape/targetstatus.qtpl:26
			qw422016.N().S(promLabelsString(promrelabel.FinalizeLabels(nil, ts.sw.Config.Labels)))
//line lib/promscrape/targetstatus.qtpl:26
			qw422016.N().S(`,`)
//line lib/promscrape/targetstatus.qtpl:26
			qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:27
			if showOriginLabels {
//line lib/promscrape/targetstatus.qtpl:27
				qw422016.N().S(`originalLabels=`)
//line lib/promscrape/targetstatus.qtpl:27
				qw422016.N().S(promLabelsString(ts.sw.Config.OriginalLabels))
//line lib/promscrape/targetstatus.qtpl:27
				qw422016.N().S(`,`)
//line lib/promscrape/targetstatus.qtpl:27
				qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:27
			}
//line lib/promscrape/targetstatus.qtpl:27
			qw422016.N().S(`scrapes_total=`)
//line lib/promscrape/targetstatus.qtpl:28
			qw422016.N().D(ts.scrapesTotal)
//line lib/promscrape/targetstatus.qtpl:28
			qw422016.N().S(`,`)
//line lib/promscrape/targetstatus.qtpl:28
			qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:28
			qw422016.N().S(`scrapes_failed=`)
//line lib/promscrape/targetstatus.qtpl:29
			qw422016.N().D(ts.scrapesFailed)
//line lib/promscrape/targetstatus.qtpl:29
			qw422016.N().S(`,`)
//line lib/promscrape/targetstatus.qtpl:29
			qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:29
			qw422016.N().S(`last_scrape=`)
//line lib/promscrape/targetstatus.qtpl:30
			qw422016.N().FPrec(ts.getDurationFromLastScrape().Seconds(), 3)
//line lib/promscrape/targetstatus.qtpl:30
			qw422016.N().S(`s ago,`)
//line lib/promscrape/targetstatus.qtpl:30
			qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:30
			qw422016.N().S(`scrape_duration=`)
//line lib/promscrape/targetstatus.qtpl:31
			qw422016.N().D(int(ts.scrapeDuration))
//line lib/promscrape/targetstatus.qtpl:31
			qw422016.N().S(`ms,`)
//line lib/promscrape/targetstatus.qtpl:31
			qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:31
			qw422016.N().S(`samples_scraped=`)
//line lib/promscrape/targetstatus.qtpl:32
			qw422016.N().D(ts.samplesScraped)
//line lib/promscrape/targetstatus.qtpl:32
			qw422016.N().S(`,`)
//line lib/promscrape/targetstatus.qtpl:32
			qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:32
			qw422016.N().S(`error=`)
//line lib/promscrape/targetstatus.qtpl:33
			if ts.err != nil {
//line lib/promscrape/targetstatus.qtpl:33
				qw422016.N().S(ts.err.Error())
//line lib/promscrape/targetstatus.qtpl:33
			}
//line lib/promscrape/targetstatus.qtpl:34
			qw422016.N().S(`
`)
//line lib/promscrape/targetstatus.qtpl:35
		}
//line lib/promscrape/targetstatus.qtpl:36
	}
//line lib/promscrape/targetstatus.qtpl:38
	for _, jobName := range emptyJobs {
//line lib/promscrape/targetstatus.qtpl:38
		qw422016.N().S(`job=`)
//line lib/promscrape/targetstatus.qtpl:39
		qw422016.N().Q(jobName)
//line lib/promscrape/targetstatus.qtpl:39
		qw422016.N().S(`(0/0 up)`)
//line lib/promscrape/targetstatus.qtpl:40
		qw422016.N().S(`
`)
//line lib/promscrape/targetstatus.qtpl:41
	}
//line lib/promscrape/targetstatus.qtpl:43
}

//line lib/promscrape/targetstatus.qtpl:43
func WriteTargetsResponsePlain(qq422016 qtio422016.Writer, jts []jobTargetsStatuses, emptyJobs []string, showOriginLabels, showOnlyUnhealthy bool, err error) {
//line lib/promscrape/targetstatus.qtpl:43
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/promscrape/targetstatus.qtpl:43
	StreamTargetsResponsePlain(qw422016, jts, emptyJobs, showOriginLabels, showOnlyUnhealthy, err)
//line lib/promscrape/targetstatus.qtpl:43
	qt422016.ReleaseWriter(qw422016)
//line lib/promscrape/targetstatus.qtpl:43
}

//line lib/promscrape/targetstatus.qtpl:43
func TargetsResponsePlain(jts []jobTargetsStatuses, emptyJobs []string, showOriginLabels, showOnlyUnhealthy bool, err error) string {
//line lib/promscrape/targetstatus.qtpl:43
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/promscrape/targetstatus.qtpl:43
	WriteTargetsResponsePlain(qb422016, jts, emptyJobs, showOriginLabels, showOnlyUnhealthy, err)
//line lib/promscrape/targetstatus.qtpl:43
	qs422016 := string(qb422016.B)
//line lib/promscrape/targetstatus.qtpl:43
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/promscrape/targetstatus.qtpl:43
	return qs422016
//line lib/promscrape/targetstatus.qtpl:43
}

//line lib/promscrape/targetstatus.qtpl:45
func StreamTargetsResponseHTML(qw422016 *qt422016.Writer, jts []jobTargetsStatuses, emptyJobs []string, showOnlyUnhealthy bool, endpointSearch, labelSearch string, err error) {
//line lib/promscrape/targetstatus.qtpl:45
	qw422016.N().S(`<!DOCTYPE html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC" crossorigin="anonymous"><title>Scrape targets</title><script>function collapse_all() {for (var i = 0; i <`)
//line lib/promscrape/targetstatus.qtpl:55
	qw422016.N().D(len(jts))
//line lib/promscrape/targetstatus.qtpl:55
	qw422016.N().S(`; i++) {let el = document.getElementById("table-" + i);if (!el) {continue;}el.style.display = 'none';}}function expand_all() {for (var i = 0; i <`)
//line lib/promscrape/targetstatus.qtpl:64
	qw422016.N().D(len(jts))
//line lib/promscrape/targetstatus.qtpl:64
	qw422016.N().S(`; i++) {let el = document.getElementById("table-" + i);if (!el) {continue;}el.style.display = 'block';}}</script></head><body><div class="navbar navbar-dark bg-dark box-shadow"><div class="d-flex justify-content-between"><a href="#" class="navbar-brand d-flex align-items-center ms-3" title="The High Performance Open Source Time Series Database &amp; Monitoring Solution "><svg xmlns="http://www.w3.org/2000/svg" id="VM_logo" viewBox="0 0 464.61 533.89" width="20" height="20" class="me-1"><defs><style>.cls-1{fill:#fff;}</style></defs><path class="cls-1" d="M459.86,467.77c9,7.67,24.12,13.49,39.3,13.69v0h1.68v0c15.18-.2,30.31-6,39.3-13.69,47.43-40.45,184.65-166.24,184.65-166.24,36.84-34.27-65.64-68.28-223.95-68.47h-1.68c-158.31.19-260.79,34.2-224,68.47C275.21,301.53,412.43,427.32,459.86,467.77Z" transform="translate(-267.7 -233.05)"/><path class="cls-1" d="M540.1,535.88c-9,7.67-24.12,13.5-39.3,13.7h-1.6c-15.18-.2-30.31-6-39.3-13.7-32.81-28-148.56-132.93-192.16-172.7v60.74c0,6.67,2.55,15.52,7.09,19.68,29.64,27.18,143.94,131.8,185.07,166.88,9,7.67,24.12,13.49,39.3,13.69v0h1.6v0c15.18-.2,30.31-6,39.3-13.69,41.13-35.08,155.43-139.7,185.07-166.88,4.54-4.16,7.09-13,7.09-19.68V363.18C688.66,403,572.91,507.9,540.1,535.88Z" transform="translate(-267.7 -233.05)"/><path class="cls-1" d="M540.1,678.64c-9,7.67-24.12,13.49-39.3,13.69v0h-1.6v0c-15.18-.2-30.31-6-39.3-13.69-32.81-28-148.56-132.94-192.16-172.7v60.73c0,6.67,2.55,15.53,7.09,19.69,29.64,27.17,143.94,131.8,185.07,166.87,9,7.67,24.12,13.5,39.3,13.7h1.6c15.18-.2,30.31-6,39.3-13.7,41.13-35.07,155.43-139.7,185.07-166.87,4.54-4.16,7.09-13,7.09-19.69V505.94C688.66,545.7,572.91,650.66,540.1,678.64Z" transform="translate(-267.7 -233.05)"/></svg><strong>VictoriaMetrics</strong></a></div></div><div class="container-fluid">`)
//line lib/promscrape/targetstatus.qtpl:84
	if err != nil {
//line lib/promscrape/targetstatus.qtpl:85
		streamerrorNotification(qw422016, err)
//line lib/promscrape/targetstatus.qtpl:86
	}
//line lib/promscrape/targetstatus.qtpl:86
	qw422016.N().S(`<div class="row"><main class="col-12"><h1>Scrape targets</h1><hr /><div class="row g-3 align-items-center mb-3"><div class="col-auto"><button type="button" class="btn`)
//line lib/promscrape/targetstatus.qtpl:93
	qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:93
	if !showOnlyUnhealthy {
//line lib/promscrape/targetstatus.qtpl:93
		qw422016.N().S(`btn-secondary`)
//line lib/promscrape/targetstatus.qtpl:93
	} else {
//line lib/promscrape/targetstatus.qtpl:93
		qw422016.N().S(`btn-success`)
//line lib/promscrape/targetstatus.qtpl:93
	}
//line lib/promscrape/targetstatus.qtpl:93
	qw422016.N().S(`" onclick="location.href='?`)
//line lib/promscrape/targetstatus.qtpl:93
	streamqueryArgs(qw422016, map[string]string{
		"show_only_unhealthy": "false",
		"endpoint_search":     endpointSearch,
		"label_search":        labelSearch,
	})
//line lib/promscrape/targetstatus.qtpl:97
	qw422016.N().S(`'">All</button></div><div class="col-auto"><button type="button" class="btn`)
//line lib/promscrape/targetstatus.qtpl:102
	qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:102
	if showOnlyUnhealthy {
//line lib/promscrape/targetstatus.qtpl:102
		qw422016.N().S(`btn-secondary`)
//line lib/promscrape/targetstatus.qtpl:102
	} else {
//line lib/promscrape/targetstatus.qtpl:102
		qw422016.N().S(`btn-danger`)
//line lib/promscrape/targetstatus.qtpl:102
	}
//line lib/promscrape/targetstatus.qtpl:102
	qw422016.N().S(`" onclick="location.href='?`)
//line lib/promscrape/targetstatus.qtpl:102
	streamqueryArgs(qw422016, map[string]string{
		"show_only_unhealthy": "true",
		"endpoint_search":     endpointSearch,
		"label_search":        labelSearch,
	})
//line lib/promscrape/targetstatus.qtpl:106
	qw422016.N().S(`'">Unhealthy</button></div><div class="col-auto"><button type="button" class="btn btn-primary" onclick="collapse_all()">Collapse all</button></div><div class="col-auto"><button type="button" class="btn btn-secondary" onclick="expand_all()">Expand all</button></div><div class="col-auto">`)
//line lib/promscrape/targetstatus.qtpl:121
	if endpointSearch == "" && labelSearch == "" {
//line lib/promscrape/targetstatus.qtpl:121
		qw422016.N().S(`<button type="button" class="btn btn-success" onclick="document.getElementById('filters').style.display='block'">Filter targets</button>`)
//line lib/promscrape/targetstatus.qtpl:125
	} else {
//line lib/promscrape/targetstatus.qtpl:125
		qw422016.N().S(`<button type="button" class="btn btn-danger" onclick="location.href='?'">Clear target filters</button>`)
//line lib/promscrape/targetstatus.qtpl:129
	}
//line lib/promscrape/targetstatus.qtpl:129
	qw422016.N().S(`</div></div><div id="filters"`)
//line lib/promscrape/targetstatus.qtpl:132
	if endpointSearch == "" && labelSearch == "" {
//line lib/promscrape/targetstatus.qtpl:132
		qw422016.N().S(`style="display:none"`)
//line lib/promscrape/targetstatus.qtpl:132
	}
//line lib/promscrape/targetstatus.qtpl:132
	qw422016.N().S(`><form class="form-horizontal"><div class="form-group mb-3"><label for="endpoint_search" class="col-sm-10 control-label">Endpoint filter (<a target="_blank" href="https://github.com/google/re2/wiki/Syntax">Regexp</a> is accepted)</label><div class="col-sm-10"><input type="text" id="endpoint_search" name="endpoint_search"placeholder="For example, 127.0.0.1" class="form-control" value="`)
//line lib/promscrape/targetstatus.qtpl:138
	qw422016.E().S(endpointSearch)
//line lib/promscrape/targetstatus.qtpl:138
	qw422016.N().S(`"/></div></div><div class="form-group mb-3"><label for="label_search" class="col-sm-10 control-label">Labels filter (<a target="_blank" href="https://prometheus.io/docs/prometheus/latest/querying/basics/#time-series-selectors">Arbitrary time series selectors</a> are accepted)</label><div class="col-sm-10"><input type="text" id="label_search" name="label_search"placeholder="For example, {instance=~'.+:9100'}" class="form-control" value="`)
//line lib/promscrape/targetstatus.qtpl:145
	qw422016.E().S(labelSearch)
//line lib/promscrape/targetstatus.qtpl:145
	qw422016.N().S(`"/></div></div><input type="hidden" name="show_only_unhealthy" value="`)
//line lib/promscrape/targetstatus.qtpl:148
	qw422016.E().V(showOnlyUnhealthy)
//line lib/promscrape/targetstatus.qtpl:148
	qw422016.N().S(`"/><button type="submit" class="btn btn-success mb-3">Submit</button></form></div><hr /><div class="row"><div class="col-12">`)
//line lib/promscrape/targetstatus.qtpl:155
	for i, js := range jts {
//line lib/promscrape/targetstatus.qtpl:156
		if showOnlyUnhealthy && js.upCount == js.targetsTotal {
//line lib/promscrape/targetstatus.qtpl:156
			continue
//line lib/promscrape/targetstatus.qtpl:156
		}
//line lib/promscrape/targetstatus.qtpl:156
		qw422016.N().S(`<div class="row mb-4"><div class="col-12"><h4><span class="me-2">`)
//line lib/promscrape/targetstatus.qtpl:160
		qw422016.E().S(js.job)
//line lib/promscrape/targetstatus.qtpl:160
		qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:160
		qw422016.N().S(`(`)
//line lib/promscrape/targetstatus.qtpl:160
		qw422016.N().D(js.upCount)
//line lib/promscrape/targetstatus.qtpl:160
		qw422016.N().S(`/`)
//line lib/promscrape/targetstatus.qtpl:160
		qw422016.N().D(js.targetsTotal)
//line lib/promscrape/targetstatus.qtpl:160
		qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:160
		qw422016.N().S(`up)</span><button type="button" class="btn btn-primary btn-sm me-1"onclick="document.getElementById('table-`)
//line lib/promscrape/targetstatus.qtpl:162
		qw422016.N().D(i)
//line lib/promscrape/targetstatus.qtpl:162
		qw422016.N().S(`').style.display='none'">collapse</button><button type="button" class="btn btn-secondary btn-sm me-1"onclick="document.getElementById('table-`)
//line lib/promscrape/targetstatus.qtpl:165
		qw422016.N().D(i)
//line lib/promscrape/targetstatus.qtpl:165
		qw422016.N().S(`').style.display='block'">expand</button></h4><div id="table-`)
//line lib/promscrape/targetstatus.qtpl:168
		qw422016.N().D(i)
//line lib/promscrape/targetstatus.qtpl:168
		qw422016.N().S(`" class="table-responsive"><table class="table table-striped table-hover table-bordered table-sm"><thead><tr><th scope="col">Endpoint</th><th scope="col">State</th><th scope="col" title="scrape target labels">Labels</th><th scope="col" title="total scrapes">Scrapes</th><th scope="col" title="total scrape errors">Errors</th><th scope="col" title="the time of the last scrape">Last Scrape</th><th scope="col" title="the duration of the last scrape">Duration</th><th scope="col" title="the number of metrics scraped during the last scrape">Samples</th><th scope="col" title="error from the last scrape (if any)">Last error</th></tr></thead><tbody class="list-`)
//line lib/promscrape/targetstatus.qtpl:183
		qw422016.N().D(i)
//line lib/promscrape/targetstatus.qtpl:183
		qw422016.N().S(`">`)
//line lib/promscrape/targetstatus.qtpl:184
		for _, ts := range js.targetsStatus {
//line lib/promscrape/targetstatus.qtpl:186
			endpoint := ts.sw.Config.ScrapeURL
			targetID := getTargetID(ts.sw)
			lastScrapeTime := ts.getDurationFromLastScrape()

//line lib/promscrape/targetstatus.qtpl:190
			if showOnlyUnhealthy && ts.up {
//line lib/promscrape/targetstatus.qtpl:190
				continue
//line lib/promscrape/targetstatus.qtpl:190
			}
//line lib/promscrape/targetstatus.qtpl:190
			qw422016.N().S(`<tr`)
//line lib/promscrape/targetstatus.qtpl:191
			if !ts.up {
//line lib/promscrape/targetstatus.qtpl:191
				qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:191
				qw422016.N().S(`class="alert alert-danger" role="alert"`)
//line lib/promscrape/targetstatus.qtpl:191
			}
//line lib/promscrape/targetstatus.qtpl:191
			qw422016.N().S(`><td class="endpoint"><a href="`)
//line lib/promscrape/targetstatus.qtpl:192
			qw422016.E().S(endpoint)
//line lib/promscrape/targetstatus.qtpl:192
			qw422016.N().S(`" target="_blank">`)
//line lib/promscrape/targetstatus.qtpl:192
			qw422016.E().S(endpoint)
//line lib/promscrape/targetstatus.qtpl:192
			qw422016.N().S(`</a> (<a href="target_response?id=`)
//line lib/promscrape/targetstatus.qtpl:193
			qw422016.E().S(targetID)
//line lib/promscrape/targetstatus.qtpl:193
			qw422016.N().S(`" target="_blank"title="click to fetch target response on behalf of the scraper">response</a>)</td><td>`)
//line lib/promscrape/targetstatus.qtpl:197
			if ts.up {
//line lib/promscrape/targetstatus.qtpl:197
				qw422016.N().S(`UP`)
//line lib/promscrape/targetstatus.qtpl:197
			} else {
//line lib/promscrape/targetstatus.qtpl:197
				qw422016.N().S(`DOWN`)
//line lib/promscrape/targetstatus.qtpl:197
			}
//line lib/promscrape/targetstatus.qtpl:197
			qw422016.N().S(`</td><td class="labels"><div title="click to show original labels"onclick="document.getElementById('original_labels_`)
//line lib/promscrape/targetstatus.qtpl:200
			qw422016.E().S(targetID)
//line lib/promscrape/targetstatus.qtpl:200
			qw422016.N().S(`').style.display='block'">`)
//line lib/promscrape/targetstatus.qtpl:201
			streamformatLabel(qw422016, promrelabel.FinalizeLabels(nil, ts.sw.Config.Labels))
//line lib/promscrape/targetstatus.qtpl:201
			qw422016.N().S(`</div><div style="display:none" id="original_labels_`)
//line lib/promscrape/targetstatus.qtpl:203
			qw422016.E().S(targetID)
//line lib/promscrape/targetstatus.qtpl:203
			qw422016.N().S(`">`)
//line lib/promscrape/targetstatus.qtpl:204
			streamformatLabel(qw422016, ts.sw.Config.OriginalLabels)
//line lib/promscrape/targetstatus.qtpl:204
			qw422016.N().S(`</div></td><td>`)
//line lib/promscrape/targetstatus.qtpl:207
			qw422016.N().D(ts.scrapesTotal)
//line lib/promscrape/targetstatus.qtpl:207
			qw422016.N().S(`</td><td>`)
//line lib/promscrape/targetstatus.qtpl:208
			qw422016.N().D(ts.scrapesFailed)
//line lib/promscrape/targetstatus.qtpl:208
			qw422016.N().S(`</td><td>`)
//line lib/promscrape/targetstatus.qtpl:210
			if lastScrapeTime < 365*24*time.Hour {
//line lib/promscrape/targetstatus.qtpl:211
				qw422016.N().FPrec(lastScrapeTime.Seconds(), 3)
//line lib/promscrape/targetstatus.qtpl:211
				qw422016.N().S(`s ago`)
//line lib/promscrape/targetstatus.qtpl:212
			} else {
//line lib/promscrape/targetstatus.qtpl:212
				qw422016.N().S(`none`)
//line lib/promscrape/targetstatus.qtpl:214
			}
//line lib/promscrape/targetstatus.qtpl:214
			qw422016.N().S(`<td>`)
//line lib/promscrape/targetstatus.qtpl:215
			qw422016.N().D(int(ts.scrapeDuration))
//line lib/promscrape/targetstatus.qtpl:215
			qw422016.N().S(`ms</td><td>`)
//line lib/promscrape/targetstatus.qtpl:216
			qw422016.N().D(ts.samplesScraped)
//line lib/promscrape/targetstatus.qtpl:216
			qw422016.N().S(`</td><td>`)
//line lib/promscrape/targetstatus.qtpl:217
			if ts.err != nil {
//line lib/promscrape/targetstatus.qtpl:217
				qw422016.E().S(ts.err.Error())
//line lib/promscrape/targetstatus.qtpl:217
			}
//line lib/promscrape/targetstatus.qtpl:217
			qw422016.N().S(`</td></tr>`)
//line lib/promscrape/targetstatus.qtpl:219
		}
//line lib/promscrape/targetstatus.qtpl:219
		qw422016.N().S(`</tbody></table></div></div></div>`)
//line lib/promscrape/targetstatus.qtpl:225
	}
//line lib/promscrape/targetstatus.qtpl:225
	qw422016.N().S(`</div></div>`)
//line lib/promscrape/targetstatus.qtpl:229
	for _, jobName := range emptyJobs {
//line lib/promscrape/targetstatus.qtpl:229
		qw422016.N().S(`<div><h4><a>`)
//line lib/promscrape/targetstatus.qtpl:232
		qw422016.E().S(jobName)
//line lib/promscrape/targetstatus.qtpl:232
		qw422016.N().S(`(0/0 up)</a></h4><table class="table table-striped table-hover table-bordered table-sm"><thead><tr><th scope="col">Endpoint</th><th scope="col">State</th><th scope="col">Labels</th><th scope="col">Last Scrape</th><th scope="col">Scrape Duration</th><th scope="col">Samples Scraped</th><th scope="col">Error</th></tr></thead></table></div>`)
//line lib/promscrape/targetstatus.qtpl:248
	}
//line lib/promscrape/targetstatus.qtpl:248
	qw422016.N().S(`</main></div></div></body></html>`)
//line lib/promscrape/targetstatus.qtpl:254
}

//line lib/promscrape/targetstatus.qtpl:254
func WriteTargetsResponseHTML(qq422016 qtio422016.Writer, jts []jobTargetsStatuses, emptyJobs []string, showOnlyUnhealthy bool, endpointSearch, labelSearch string, err error) {
//line lib/promscrape/targetstatus.qtpl:254
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/promscrape/targetstatus.qtpl:254
	StreamTargetsResponseHTML(qw422016, jts, emptyJobs, showOnlyUnhealthy, endpointSearch, labelSearch, err)
//line lib/promscrape/targetstatus.qtpl:254
	qt422016.ReleaseWriter(qw422016)
//line lib/promscrape/targetstatus.qtpl:254
}

//line lib/promscrape/targetstatus.qtpl:254
func TargetsResponseHTML(jts []jobTargetsStatuses, emptyJobs []string, showOnlyUnhealthy bool, endpointSearch, labelSearch string, err error) string {
//line lib/promscrape/targetstatus.qtpl:254
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/promscrape/targetstatus.qtpl:254
	WriteTargetsResponseHTML(qb422016, jts, emptyJobs, showOnlyUnhealthy, endpointSearch, labelSearch, err)
//line lib/promscrape/targetstatus.qtpl:254
	qs422016 := string(qb422016.B)
//line lib/promscrape/targetstatus.qtpl:254
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/promscrape/targetstatus.qtpl:254
	return qs422016
//line lib/promscrape/targetstatus.qtpl:254
}

//line lib/promscrape/targetstatus.qtpl:256
func streamqueryArgs(qw422016 *qt422016.Writer, m map[string]string) {
//line lib/promscrape/targetstatus.qtpl:258
	qa := make(url.Values, len(m))
	for k, v := range m {
		qa[k] = []string{v}
	}

//line lib/promscrape/targetstatus.qtpl:263
	qw422016.E().S(qa.Encode())
//line lib/promscrape/targetstatus.qtpl:264
}

//line lib/promscrape/targetstatus.qtpl:264
func writequeryArgs(qq422016 qtio422016.Writer, m map[string]string) {
//line lib/promscrape/targetstatus.qtpl:264
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/promscrape/targetstatus.qtpl:264
	streamqueryArgs(qw422016, m)
//line lib/promscrape/targetstatus.qtpl:264
	qt422016.ReleaseWriter(qw422016)
//line lib/promscrape/targetstatus.qtpl:264
}

//line lib/promscrape/targetstatus.qtpl:264
func queryArgs(m map[string]string) string {
//line lib/promscrape/targetstatus.qtpl:264
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/promscrape/targetstatus.qtpl:264
	writequeryArgs(qb422016, m)
//line lib/promscrape/targetstatus.qtpl:264
	qs422016 := string(qb422016.B)
//line lib/promscrape/targetstatus.qtpl:264
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/promscrape/targetstatus.qtpl:264
	return qs422016
//line lib/promscrape/targetstatus.qtpl:264
}

//line lib/promscrape/targetstatus.qtpl:266
func streamformatLabel(qw422016 *qt422016.Writer, labels []prompbmarshal.Label) {
//line lib/promscrape/targetstatus.qtpl:266
	qw422016.N().S(`{`)
//line lib/promscrape/targetstatus.qtpl:268
	for i, label := range labels {
//line lib/promscrape/targetstatus.qtpl:269
		qw422016.E().S(label.Name)
//line lib/promscrape/targetstatus.qtpl:269
		qw422016.N().S(`=`)
//line lib/promscrape/targetstatus.qtpl:269
		qw422016.E().Q(label.Value)
//line lib/promscrape/targetstatus.qtpl:270
		if i+1 < len(labels) {
//line lib/promscrape/targetstatus.qtpl:270
			qw422016.N().S(`,`)
//line lib/promscrape/targetstatus.qtpl:270
			qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:270
		}
//line lib/promscrape/targetstatus.qtpl:271
	}
//line lib/promscrape/targetstatus.qtpl:271
	qw422016.N().S(`}`)
//line lib/promscrape/targetstatus.qtpl:273
}

//line lib/promscrape/targetstatus.qtpl:273
func writeformatLabel(qq422016 qtio422016.Writer, labels []prompbmarshal.Label) {
//line lib/promscrape/targetstatus.qtpl:273
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/promscrape/targetstatus.qtpl:273
	streamformatLabel(qw422016, labels)
//line lib/promscrape/targetstatus.qtpl:273
	qt422016.ReleaseWriter(qw422016)
//line lib/promscrape/targetstatus.qtpl:273
}

//line lib/promscrape/targetstatus.qtpl:273
func formatLabel(labels []prompbmarshal.Label) string {
//line lib/promscrape/targetstatus.qtpl:273
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/promscrape/targetstatus.qtpl:273
	writeformatLabel(qb422016, labels)
//line lib/promscrape/targetstatus.qtpl:273
	qs422016 := string(qb422016.B)
//line lib/promscrape/targetstatus.qtpl:273
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/promscrape/targetstatus.qtpl:273
	return qs422016
//line lib/promscrape/targetstatus.qtpl:273
}

//line lib/promscrape/targetstatus.qtpl:275
func streamerrorNotification(qw422016 *qt422016.Writer, err error) {
//line lib/promscrape/targetstatus.qtpl:275
	qw422016.N().S(`<div class="alert alert-danger d-flex align-items-center" role="alert"><svg class="bi flex-shrink-0 me-2" width="24" height="24" role="img" aria-label="Danger:"><use xlink:href="#exclamation-triangle-fill"/></svg><div>`)
//line lib/promscrape/targetstatus.qtpl:280
	qw422016.E().S(err.Error())
//line lib/promscrape/targetstatus.qtpl:280
	qw422016.N().S(`</div></div>`)
//line lib/promscrape/targetstatus.qtpl:283
}

//line lib/promscrape/targetstatus.qtpl:283
func writeerrorNotification(qq422016 qtio422016.Writer, err error) {
//line lib/promscrape/targetstatus.qtpl:283
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/promscrape/targetstatus.qtpl:283
	streamerrorNotification(qw422016, err)
//line lib/promscrape/targetstatus.qtpl:283
	qt422016.ReleaseWriter(qw422016)
//line lib/promscrape/targetstatus.qtpl:283
}

//line lib/promscrape/targetstatus.qtpl:283
func errorNotification(err error) string {
//line lib/promscrape/targetstatus.qtpl:283
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/promscrape/targetstatus.qtpl:283
	writeerrorNotification(qb422016, err)
//line lib/promscrape/targetstatus.qtpl:283
	qs422016 := string(qb422016.B)
//line lib/promscrape/targetstatus.qtpl:283
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/promscrape/targetstatus.qtpl:283
	return qs422016
//line lib/promscrape/targetstatus.qtpl:283
}
