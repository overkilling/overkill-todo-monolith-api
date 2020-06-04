package http

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Router provides REST endpoint routing for the Todo API
func Router(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))

	r.Get("/health", NewHealthcheckHandler(func() bool { return db.Ping() == nil }))
	r.Get("/todos", NewTodosHandler())

	return r
}
