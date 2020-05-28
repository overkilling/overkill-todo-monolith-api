package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/overkilling/overkill-todo-monolith-api/postgres"

	_ "github.com/lib/pq"
)

type serviceUpCheck func() bool

func healthcheckHandler(dbServiceAlive serviceUpCheck) http.HandlerFunc {
	type HealthcheckResponse struct {
		Status string `json:"status"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		status := "fail"
		if dbServiceAlive() {
			status = "ok"
		}
		response := HealthcheckResponse{status}
		responseBytes, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseBytes)
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
