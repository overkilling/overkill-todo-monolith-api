package postgres

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Only relevant for migrations
)

// MigrateDB applies the database migrations into the supplied DB
func MigrateDB(db *sql.DB, migrationsPath string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	migrator, err := migrate.NewWithDatabaseInstance(migrationsPath, "postgres", driver)
	if err != nil {
		return err
	}

	err = migrator.Up()
	if err == migrate.ErrNoChange || err == migrate.ErrInvalidVersion {
		return nil
	}
	return err
}
