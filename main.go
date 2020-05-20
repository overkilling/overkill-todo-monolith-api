package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/overkilling/overkill-todo-monolith-api/postgres"

	_ "github.com/lib/pq"
)

type serviceUpCheck func() bool

func healthcheckHandler(dbServiceAlive serviceUpCheck) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result := "db down"
		if dbServiceAlive() {
			result = "ok"
		}
		w.Write([]byte(result))
	}
}

func router(dbName string) http.Handler {
	db := postgres.NewDb(
		postgres.ConfigOption{Option: "dbname", Value: dbName},
		postgres.ConfigOption{Option: "host", Value: "localhost"},
		postgres.ConfigOption{Option: "port", Value: "5432"},
		postgres.ConfigOption{Option: "user", Value: "postgres"},
		postgres.ConfigOption{Option: "password", Value: "postgres"},
		postgres.ConfigOption{Option: "sslmode", Value: "disable"},
	)
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))

	r.Get("/health", healthcheckHandler(db.Alive))

	return r
}

func main() {
	http.ListenAndServe(":3000", router("todo"))
}
