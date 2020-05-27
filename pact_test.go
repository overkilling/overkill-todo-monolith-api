// +build pact

package main

import (
	"fmt"
	"net/http"
	"testing"

	_ "github.com/lib/pq"
	"github.com/overkilling/overkill-todo-monolith-api/postgres"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
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

func TestPactProvider(t *testing.T) {
	go startProvider()

	pact := createPact()
	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: "http://127.0.0.1:3000",
		PactURLs:        []string{"https://raw.githubusercontent.com/overkilling/overkill-todo-infrastructure/master/pacts/spa-api.json"},
	})

	if err != nil {
		t.Log("Pact test failed")
	}
}

func startProvider() {
	recreateTestDb("todo_test")
	db, err := postgres.NewDb(append(dbBaseConfig, postgres.DbName("todo_test"))...)
	if err != nil {
		panic(err)
	}

	err = http.ListenAndServe(":3000", router(db))
	if err != nil {
		panic(err)
	}
}

func createPact() dsl.Pact {
	return dsl.Pact{
		Provider: "API",
	}
}
