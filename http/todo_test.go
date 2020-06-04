package http_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	todo "github.com/overkilling/overkill-todo-monolith-api"
	todoHttp "github.com/overkilling/overkill-todo-monolith-api/http"
	"github.com/stretchr/testify/assert"
)

func TestGetTodosHandler(t *testing.T) {
	fecher := func() ([]todo.Todo, error) {
		return []todo.Todo{
			{Todo: "Some todo"},
			{Todo: "Another todo"},
			{Todo: "Yet another todo"},
		}, nil
	}

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:3000/todos", nil)

	todoHttp.NewTodosHandler(fecher)(res, req)

	content, _ := ioutil.ReadAll(res.Body)
	assert.JSONEq(t, `
		[
			{"todo": "Some todo"},
			{"todo": "Another todo"},
			{"todo": "Yet another todo"}
		]
	`, string(content))
}
