// +build integration

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/lib/pq"
	"github.com/overkilling/overkill-todo-monolith-api/postgres"
	"github.com/stretchr/testify/assert"
)

var dbBaseConfig = []postgres.ConfigOption{
	postgres.Credentials("postgres", "postgres"),
	postgres.HostAndPort("localhost", 5432),
	postgres.SslDisabled(),
}

func recreateTestDb(dbName string) {
	var err error

	db, err := postgres.NewDb(dbBaseConfig...)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", dbName))
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
	if err != nil {
		panic(err)
	}
}

func TestIntegrationRouter(t *testing.T) {
	recreateTestDb("todo_test")
	db, err := postgres.NewDb(append(dbBaseConfig, postgres.DbName("todo_test"))...)
	if err != nil {
		panic(err)
	}

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:3000/health", nil)

	router(db).ServeHTTP(res, req)

	content, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, "ok", string(content))
}
