package todo

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	todoHttp "github.com/overkilling/overkill-todo-monolith-api/http"
)

func getTodosHandler() http.HandlerFunc {
	type Todo struct {
		Todo string `json:"todo"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		todos := [1]Todo{{Todo: "Some task"}}
		responseBytes, _ := json.Marshal(todos)
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseBytes)
	}
}

// Router provides REST endpoint routing for the Todo API
func Router(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))

	r.Get("/health", todoHttp.NewHealthcheckHandler(func() bool { return db.Ping() == nil }))
	r.Get("/todos", getTodosHandler())

	return r
}
