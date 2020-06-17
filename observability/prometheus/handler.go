package prometheus

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var instrumentationLabels = []string{"code", "method"}

// Handler wraps Prometheus package handler
func Handler() http.Handler {
	return promhttp.Handler()
}

// Instrument a HTTP handler with counter, duration and inflight metrics.
func Instrument(label string, handler http.Handler) http.Handler {
	requestCounter := promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_handler_requests_total", label),
			Help: fmt.Sprintf("Total number of %s requests by HTTP code and method", label),
		},
		instrumentationLabels,
	)
	inFlightGauge := promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: fmt.Sprintf("%s_handler_requests_in_flight", label),
			Help: fmt.Sprintf("Number of %s requests currently in flight", label),
		},
	)
	durationHistogram := promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: fmt.Sprintf("%s_handler_requests_duration", label),
			Help: fmt.Sprintf("Histogram of %s request durations in seconds", label),
		},
		instrumentationLabels,
	)

	return promhttp.InstrumentHandlerCounter(requestCounter,
		promhttp.InstrumentHandlerInFlight(inFlightGauge,
			promhttp.InstrumentHandlerDuration(durationHistogram,
				handler)))
}
