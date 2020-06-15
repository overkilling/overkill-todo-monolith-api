package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	t.Run("default values", func(t *testing.T) {
		config := loadConfig()

		assert.Equal(t, "localhost", config.db.host)
		assert.Equal(t, 5432, config.db.port)
		assert.Equal(t, "todo", config.db.database)
		assert.Equal(t, "postgres", config.db.username)
		assert.Equal(t, "postgres", config.db.password)
	})
	t.Run("from environment variables", func(t *testing.T) {
		os.Setenv("DB_HOST", "hostfromenv")
		os.Setenv("DB_PORT", "1234")
		os.Setenv("DB_DATABASE", "dbfromenv")
		os.Setenv("DB_USERNAME", "userfromenv")
		os.Setenv("DB_PASSWORD", "passfromenv")

		config := loadConfig()

		assert.Equal(t, "hostfromenv", config.db.host)
		assert.Equal(t, 1234, config.db.port)
		assert.Equal(t, "dbfromenv", config.db.database)
		assert.Equal(t, "userfromenv", config.db.username)
		assert.Equal(t, "passfromenv", config.db.password)
	})
}
