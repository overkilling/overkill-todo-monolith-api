package main

import (
	"database/sql"
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

func router(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))

	r.Get("/health", healthcheckHandler(func() bool { return db.Ping() == nil }))

	return r
}

func main() {
	db, err := postgres.NewDb(
		postgres.DbName("todo"),
		postgres.Credentials("postgres", "postgres"),
		postgres.HostAndPort("localhost", 5432),
		postgres.SslDisabled(),
	)
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":3000", router(db))
}
