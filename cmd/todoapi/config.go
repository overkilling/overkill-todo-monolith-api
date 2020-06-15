package main

import (
	"strings"

	"github.com/spf13/viper"
)

func configureViper() {
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", 5432)
	viper.SetDefault("db.database", "todo")
	viper.SetDefault("db.username", "postgres")
	viper.SetDefault("db.password", "postgres")
	viper.SetDefault("log.pretty", false)

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.AddConfigPath(".")

	viper.ReadInConfig()
}

type config struct {
	db  dbConfig
	log logConfig
}

type dbConfig struct {
	host     string
	port     int
	database string
	username string
	password string
}

type logConfig struct {
	pretty bool
}

func loadConfig() config {
	configureViper()

	return config{
		db: dbConfig{
			host:     viper.GetString("db.host"),
			port:     viper.GetInt("db.port"),
			database: viper.GetString("db.database"),
			username: viper.GetString("db.username"),
			password: viper.GetString("db.password"),
		},
		log: logConfig{
			pretty: viper.GetBool("log.pretty"),
		},
	}
}
