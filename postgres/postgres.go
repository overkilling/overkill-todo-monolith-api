package postgres

import (
	"database/sql"
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

// NewDb returns a new database access object for a postgres database, from
// a set of configuration options.
func NewDb(configs ...configOption) (*Db, error) {
	db, err := sql.Open("postgres", toConnString(configs))
	return &Db{sqlDb: db}, err
}
