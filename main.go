package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	_ "github.com/lib/pq"
)

type database interface {
	Alive() bool
}

type wrappedSQLDb struct {
	sqlDb *sql.DB
}

func (db wrappedSQLDb) Alive() bool {
	return db.sqlDb.Ping() == nil
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "todo"
)

func newDb() database {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return wrappedSQLDb{sqlDb: db}
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	result := "ok"
	if !newDb().Alive() {
		result = "db down"
	}
	w.Write([]byte(result))
}

func router() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))

	r.Get("/health", healthcheckHandler)

	return r
}

func main() {
	http.ListenAndServe(":3000", router())
}
