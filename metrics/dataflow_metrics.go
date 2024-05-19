package metrics

import (
	"github.com/fdataflow/fcommon"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func RunMetricsService(serverAddr string) error {
	http.Handle(fcommon.METRICS_ROUTE, promhttp.Handler())
	return http.ListenAndServe(serverAddr, nil)
}
