// Package cutil - Content managed by Project Forge, see [projectforge.md] for details.
package cutil

import (
	"net/http"
	"slices"

	"github.com/CAFxX/httpcompression"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/kyleu/rituals/app/lib/telemetry/httpmetrics"
	"github.com/kyleu/rituals/app/util"
)

var AppRoutesList = map[string][]string{}

func WireRouter(r *mux.Router, logger util.Logger) (http.Handler, error) {
	p := httpmetrics.NewMetrics(util.AppKey, logger)
	r.Handle(p.MetricsPath, promhttp.Handler()).Methods(http.MethodGet)

	var ret http.Handler = r
	includeCompression := util.GetEnvBool("compression_enabled", false)
	if includeCompression {
		compress, _ := httpcompression.DefaultAdapter()
		ret = compress(ret)
	}
	return p.WrapHandler(r), nil
}

func AddRoute(method string, path string) {
	curr := AppRoutesList[method]
	if !slices.Contains(curr, path) {
		curr = append(curr, path)
		slices.Sort(curr)
		AppRoutesList[method] = curr
	}
}
