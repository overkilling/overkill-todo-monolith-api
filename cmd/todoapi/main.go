package main

import (
	"os"
	"strings"

	_ "github.com/lib/pq"
	"github.com/overkilling/overkill-todo-monolith-api/http"
	"github.com/overkilling/overkill-todo-monolith-api/postgres"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

func configureViper() {
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", 5432)
	viper.SetDefault("db.database", "todo")
	viper.SetDefault("db.username", "postgres")
	viper.SetDefault("db.password", "postgres")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.AddConfigPath(".")

	viper.ReadInConfig()
}

func main() {
	configureViper()
	log := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("service", "api").
		Logger()

	db, err := postgres.NewDb(
		postgres.DbName(viper.GetString("db.database")),
		postgres.Credentials(viper.GetString("db.username"), viper.GetString("db.password")),
		postgres.HostAndPort(viper.GetString("db.host"), viper.GetInt("db.port")),
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
