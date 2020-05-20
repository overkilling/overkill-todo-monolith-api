// +build integration

package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func recreateTestDb(dbName string) {
	var err error

	db, err := sql.Open("postgres", "user=postgres password=postgres sslmode=disable")
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

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:3000/health", nil)

	router("todo_test").ServeHTTP(res, req)

	content, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, "ok", string(content))
}
