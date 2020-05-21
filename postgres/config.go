package postgres

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

// NewDb returns a new database access object for a postgres database, from
// a set of configuration options.
func NewDb(configs ...configOption) (*sql.DB, error) {
	return sql.Open("postgres", toConnString(configs))
}

type configOption func(configBuilder)

type configBuilder func(string, string)

type connStrConfig struct {
	builder strings.Builder
}

func (c *connStrConfig) append(key, value string) {
	fmt.Fprintf(&c.builder, "%s=%s ", key, value)
}

func toConnString(configs []configOption) string {
	connStrBuilder := connStrConfig{}
	for _, config := range configs {
		config(connStrBuilder.append)
	}
	return connStrBuilder.builder.String()
}

// DbName configures postgres with a database name
func DbName(dbName string) func(configBuilder) {
	return func(config configBuilder) {
		config("dbname", dbName)
	}
}

// Credentials configures postgres with username and password
func Credentials(username, password string) func(configBuilder) {
	return func(config configBuilder) {
		config("user", username)
		config("password", password)
	}
}

// HostAndPort configures postgres with hostname and port number
func HostAndPort(hostname string, port int) func(configBuilder) {
	return func(config configBuilder) {
		config("host", hostname)
		config("port", strconv.Itoa(port))
	}
}

// SslDisabled configures postgres to not enable SSL
func SslDisabled() func(configBuilder) {
	return func(config configBuilder) {
		config("sslmode", "disable")
	}
}
