package prometheus

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Handler wraps Prometheus package handler
func Handler() http.Handler {
	return promhttp.Handler()
}
