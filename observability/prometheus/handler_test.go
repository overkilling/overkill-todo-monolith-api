package prometheus_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/overkilling/overkill-todo-monolith-api/observability/prometheus"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "http://localhost:3000/metrics", nil)
	assert.NoError(t, err, "failed to create request")

	prometheus.Handler().ServeHTTP(res, req)

	content, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err, "failed to read response's body")
	assert.Contains(t, string(content), "promhttp_metric_handler_requests_total")
}
