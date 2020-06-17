package prometheus_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/overkilling/overkill-todo-monolith-api/observability/prometheus"
	"github.com/stretchr/testify/assert"
)

func getPrometheusMetrics(t *testing.T) string {
	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "http://localhost:3000/metrics", nil)
	assert.NoError(t, err, "failed to create request")
	prometheus.Handler().ServeHTTP(res, req)

	content, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err, "failed to read response's body")
	return string(content)
}

func TestInstrumentation(t *testing.T) {
	res := httptest.NewRecorder()
	testreq, _ := http.NewRequest("GET", "http://localhost:3000/test", nil)
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Test handler"))
	}
	prometheus.InstrumentHandler("test", http.HandlerFunc(handler)).ServeHTTP(res, testreq)

	content := getPrometheusMetrics(t)
	assert.Contains(t, content, `test_handler_requests_total{code="200",method="get"} 1`)
	assert.Contains(t, content, `test_handler_requests_in_flight 0`)
	assert.Contains(t, content, `test_handler_requests_duration_count{code="200",method="get"} 1`)
}
