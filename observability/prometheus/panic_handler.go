package prometheus

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// PanicHandler for counting number of panicked requests
type PanicHandler struct {
	panicCounter *prometheus.CounterVec
}

// HandlePanic increments the panicked requests counter for the
// given request's path and method
func (handler *PanicHandler) HandlePanic(req *http.Request, panicReason interface{}) {
	labels := prometheus.Labels{"path": req.URL.Path, "method": req.Method}
	handler.panicCounter.With(labels).Inc()
}

// NewPanicHandler creates a handler for a panicked requests counter.
func NewPanicHandler() *PanicHandler {
	panicCounter := promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "panicked_requests_total",
			Help: "Total number of panicked requests by HTTP path and method",
		},
		[]string{"path", "method"},
	)
	return &PanicHandler{panicCounter}
}
