package main

import (
	"io"
	"os"

	_ "github.com/lib/pq"
	"github.com/overkilling/overkill-todo-monolith-api/http"
	"github.com/overkilling/overkill-todo-monolith-api/postgres"
	"github.com/rs/zerolog"
)

func main() {
	config := loadConfig()

	var logOutput io.Writer
	logOutput = os.Stdout
	if config.log.pretty {
		logOutput = zerolog.ConsoleWriter{Out: os.Stdout}
	}
	log := zerolog.New(logOutput).With().
		Timestamp().
		Str("service", "api").
		Logger()

	db, err := postgres.NewDb(
		postgres.DbName(config.db.database),
		postgres.Credentials(config.db.username, config.db.password),
		postgres.HostAndPort(config.db.host, config.db.port),
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
