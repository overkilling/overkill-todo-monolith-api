package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
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
func NewRouter(endpoints Endpoints, log zerolog.Logger) *Router {
	mux := chi.NewRouter()

	mux.Use(hlog.NewHandler(log))
	mux.Use(hlog.MethodHandler("http_method"))
	mux.Use(hlog.URLHandler("url"))
	mux.Use(hlog.RequestIDHandler("request_id", "Request-ID"))
	mux.Use(hlog.AccessHandler(accessHandlerLogging))
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

// Serve a HTTP request
func (router *Router) Serve(w http.ResponseWriter, r *http.Request) {
	router.mux.ServeHTTP(w, r)
}

func accessHandlerLogging(r *http.Request, status, size int, duration time.Duration) {
	hlog.FromRequest(r).Info().
		Int("http_status", status).
		Int("response_size", size).
		Send()
}
