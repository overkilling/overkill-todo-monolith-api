package postgres

import (
	"database/sql"
	"fmt"
	"strings"
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

// ConfigOption represents a postgres configuration option
type ConfigOption struct {
	Option string
	Value  interface{}
}

// NewDb returns a new database access object for a postgres database, from
// a set of configuration options.
func NewDb(configs ...ConfigOption) *Db {
	fmt.Println(toConnString(configs))
	db, err := sql.Open("postgres", toConnString(configs))
	if err != nil {
		panic(err)
	}

	return &Db{sqlDb: db}
}

func toConnString(configs []ConfigOption) string {
	var connStringBuilder strings.Builder
	for _, config := range configs {
		fmt.Fprintf(&connStringBuilder, "%s=%s ", config.Option, config.Value)
	}
	return connStringBuilder.String()
}
