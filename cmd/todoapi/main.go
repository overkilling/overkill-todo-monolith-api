package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	todo "github.com/overkilling/overkill-todo-monolith-api"
	"github.com/overkilling/overkill-todo-monolith-api/postgres"
)

func main() {
	dbHost := os.Getenv("DB_HOST")

	db, err := postgres.NewDb(
		postgres.DbName("todo"),
		postgres.Credentials("postgres", "postgres"),
		postgres.HostAndPort(dbHost, 5432),
		postgres.SslDisabled(),
	)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Applying migrations...")
	err = postgres.MigrateDB(db, "file://./postgres/migrations")
	if err != nil {
		panic(err)
	}
	fmt.Println("Migrations done")

	fmt.Println("Starting server on port 3000")
	err = http.ListenAndServe(":3000", todo.Router(db))
	if err != nil {
		panic(err)
	}
}
