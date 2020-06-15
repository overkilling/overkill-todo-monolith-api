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

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.AddConfigPath(".")

	viper.ReadInConfig()
}

type config struct {
	db dbConfig
}

type dbConfig struct {
	host     string
	port     int
	database string
	username string
	password string
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
	}
}
