package http_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	todoHttp "github.com/overkilling/overkill-todo-monolith-api/http"
	"github.com/stretchr/testify/assert"
)

func TestGetTodosHandler(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:3000/todos", nil)

	todoHttp.NewTodosHandler()(res, req)

	content, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, "[{\"todo\":\"Some task\"}]", string(content))
}
