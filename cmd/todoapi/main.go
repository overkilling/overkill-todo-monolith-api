package main

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/overkilling/overkill-todo-monolith-api/http"
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

	endpoints := http.Endpoints{
		Healthcheck: http.NewHealthcheckHandler(func() bool { return db.Ping() == nil }),
		Todos:       http.NewTodosHandler(),
	}
	fmt.Println("Starting server on port 3000")
	err = http.NewRouter(endpoints).ServeOn(3000)
	if err != nil {
		panic(err)
	}
}
