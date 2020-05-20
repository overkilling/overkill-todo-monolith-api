package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/overkilling/overkill-todo-monolith-api/postgres"

	_ "github.com/lib/pq"
)

type dbAlive interface {
	Alive() bool
}

func healthcheckHandler(db dbAlive) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result := "db down"
		if db.Alive() {
			result = "ok"
		}
		w.Write([]byte(result))
	}
}

func router(dbName string) http.Handler {
	db := postgres.NewDb(dbName)
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))

	r.Get("/health", healthcheckHandler(db))

	return r
}

func main() {
	http.ListenAndServe(":3000", router("todo"))
}
