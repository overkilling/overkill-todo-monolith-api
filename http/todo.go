package http

import (
	"encoding/json"
	"net/http"

	todo "github.com/overkilling/overkill-todo-monolith-api"
)

// NewTodosHandler creates a HTTP handler for serving a list of todos.
func NewTodosHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos := [1]todo.Todo{{Todo: "Some task"}}
		responseBytes, _ := json.Marshal(todos)
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseBytes)
	}
}
