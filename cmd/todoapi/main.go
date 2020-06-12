package main

import (
	"os"

	_ "github.com/lib/pq"
	"github.com/overkilling/overkill-todo-monolith-api/http"
	"github.com/overkilling/overkill-todo-monolith-api/postgres"
	"github.com/rs/zerolog"
)

func main() {
	log := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("service", "api").
		Logger()
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

	log.Info().Str("type", "migrations").Msg("Applying migrations...")
	err = postgres.MigrateDB(db, "file://./postgres/migrations")
	if err != nil {
		log.Fatal().Str("type", "migrations").Err(err).Msg("Failed to apply migrations")
	}
	log.Info().Str("type", "migrations").Msg("Migrations done")
	todosRepository := postgres.NewTodosRepository(db)

	endpoints := http.Endpoints{
		Healthcheck: http.NewHealthcheckHandler(func() bool { return db.Ping() == nil }),
		Todos:       http.NewTodosHandler(todosRepository.GetAll),
	}
	log.Info().Str("type", "startup").Msg("Starting server on port 3000")
	err = http.NewRouter(endpoints, log).ServeOn(3000)
	if err != nil {
		panic(err)
	}
}
