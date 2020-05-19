package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func healthcheckHander(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func router() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))

	r.Get("/health", healthcheckHander)

	return r
}

func main() {
	http.ListenAndServe(":3000", router())
}
