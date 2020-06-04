package http

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Router represents a Todo's HTTP router, with the appropriate endpoints
// configured.
type Router struct {
	mux *chi.Mux
}

// NewRouter provides REST endpoint routing for the Todo API
func NewRouter(db *sql.DB) *Router {
	mux := chi.NewRouter()

	mux.Use(middleware.RequestID)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Heartbeat("/ping"))

	mux.Get("/health", NewHealthcheckHandler(func() bool { return db.Ping() == nil }))
	mux.Get("/todos", NewTodosHandler())

	return &Router{mux}
}

// ServeOn serves the router endpoints on the supplied port.
func (router *Router) ServeOn(port int) error {
	return http.ListenAndServe(fmt.Sprintf(":%d", port), router.mux)
}
