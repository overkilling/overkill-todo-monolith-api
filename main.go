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
)

func newDb(dbname string) database {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return wrappedSQLDb{sqlDb: db}
}

func healthcheckHandler(db database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result := "db down"
		if db.Alive() {
			result = "ok"
		}
		w.Write([]byte(result))
	}
}

func router(dbName string) http.Handler {
	db := newDb(dbName)
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
