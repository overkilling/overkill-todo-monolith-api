package http

import (
	"encoding/json"
	"net/http"

	todo "github.com/overkilling/overkill-todo-monolith-api"
)

type fetchTodos func() ([]todo.Todo, error)

// NewTodosHandler creates a HTTP handler for serving a list of todos.
func NewTodosHandler(fetch fetchTodos) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos, _ := fetch()
		responseBytes, _ := json.Marshal(todos)
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseBytes)
	}
}
