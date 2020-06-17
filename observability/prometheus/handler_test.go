package prometheus_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/overkilling/overkill-todo-monolith-api/observability/prometheus"
	"github.com/stretchr/testify/assert"
)

var res = httptest.NewRecorder()

func TestInstrumentation(t *testing.T) {
	testreq, _ := http.NewRequest("GET", "http://localhost:3000/test", nil)
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Test handler"))
	}
	prometheus.Instrument("test", http.HandlerFunc(handler)).ServeHTTP(res, testreq)

	req, err := http.NewRequest("GET", "http://localhost:3000/metrics", nil)
	assert.NoError(t, err, "failed to create request")
	prometheus.Handler().ServeHTTP(res, req)

	content, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err, "failed to read response's body")
	assert.Contains(t, string(content), `test_handler_requests_total{code="200",method="get"} 1`)
	assert.Contains(t, string(content), `test_handler_requests_in_flight 0`)
	assert.Contains(t, string(content), `test_handler_requests_duration_count{code="200",method="get"} 1`)
}
