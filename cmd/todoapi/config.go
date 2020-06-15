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
