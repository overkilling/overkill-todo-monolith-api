package postgres

import (
	"database/sql"
	"fmt"
)

// Db represents a postgres database handler, exposing functions
// to query the database
type Db struct {
	sqlDb *sql.DB
}

// Alive checks if the database connection is alive
func (db *Db) Alive() bool {
	return db.sqlDb.Ping() == nil
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
)

// NewDb returns a new database access object for a postgres database
func NewDb(dbname string) *Db {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return &Db{sqlDb: db}
}
