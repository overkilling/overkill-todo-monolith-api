package prometheus_test

import (
	"net/http"
	"testing"

	"github.com/overkilling/overkill-todo-monolith-api/observability/prometheus"
	"github.com/stretchr/testify/assert"
)

func TestPanicHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost:3000/some-endpoint", nil)

	prometheus.NewPanicHandler().HandlePanic(req, "some reason")

	content := getPrometheusMetrics(t)
	assert.Contains(t, content, `panicked_requests_total{method="GET",path="/some-endpoint"} 1`)
}
