package http

import (
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

// Endpoints specifies the Todo API endpoint configurations as
// a list of http.Handlers.
type Endpoints struct {
	Healthcheck http.HandlerFunc
	Todos       http.HandlerFunc
}

// NewRouter provides REST endpoint routing for the Todo API
func NewRouter(endpoints Endpoints) *Router {
	mux := chi.NewRouter()

	mux.Use(middleware.RequestID)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Heartbeat("/ping"))

	mux.Get("/health", endpoints.Healthcheck)
	mux.Get("/todos", endpoints.Todos)

	return &Router{mux}
}

// ServeOn serves the router endpoints on the supplied port.
func (router *Router) ServeOn(port int) error {
	return http.ListenAndServe(fmt.Sprintf(":%d", port), router.mux)
}
