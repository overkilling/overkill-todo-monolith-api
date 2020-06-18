package main

import (
	"os"

	_ "github.com/lib/pq"
	"github.com/overkilling/overkill-todo-monolith-api/http"
	"github.com/overkilling/overkill-todo-monolith-api/observability/prometheus"
	"github.com/overkilling/overkill-todo-monolith-api/observability/zerolog"
	"github.com/overkilling/overkill-todo-monolith-api/postgres"
)

func main() {
	config := loadConfig()

	log := zerolog.NewLogger(os.Stdout, config.log.pretty)
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
	todosRepository := prometheus.InstrumentTodosRepository(
		postgres.NewTodosRepository(db))

	endpoints := http.Endpoints{
		Metrics:     prometheus.Handler().ServeHTTP,
		Healthcheck: prometheus.InstrumentHandler("healthcheck", http.NewHealthcheckHandler(func() bool { return db.Ping() == nil })).ServeHTTP,
		Todos:       prometheus.InstrumentHandler("todos", http.NewTodosHandler(todosRepository.GetAll)).ServeHTTP,
	}
	log.Info().Str("type", "startup").Msg("Starting server on port 3000")
	err = http.NewRouter(endpoints, log, zerolog.NewPanicHandler(), prometheus.NewPanicHandler()).ServeOn(3000)
	if err != nil {
		panic(err)
	}
}
