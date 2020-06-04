package todo

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTodosHandler(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:3000/todos", nil)

	getTodosHandler()(res, req)

	content, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, "[{\"todo\":\"Some task\"}]", string(content))
}
