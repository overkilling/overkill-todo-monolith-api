package pact

import (
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/overkilling/overkill-todo-monolith-api/http"
	"github.com/overkilling/overkill-todo-monolith-api/postgres"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func TestPactProvider(t *testing.T) {
	pactURL := os.Getenv("PACT_URL")
	if pactURL == "" {
		t.Skip("set PACT_URL to run this test")
	}

	go startProvider()

	pact := createPact()
	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: "http://127.0.0.1:3000",
		PactURLs:        []string{pactURL},
	})

	if err != nil {
		t.Log("Pact test failed")
	}
}

func startProvider() {
	db, err := postgres.NewDb(
		postgres.DbName("todo"),
		postgres.Credentials("postgres", "postgres"),
		postgres.HostAndPort(os.Getenv("DB_HOST"), 5432),
		postgres.SslDisabled(),
	)
	if err != nil {
		panic(err)
	}

	err = postgres.MigrateDB(db, "file://../postgres/migrations")
	if err != nil {
		panic(err)
	}

	endpoints := http.Endpoints{
		Healthcheck: http.NewHealthcheckHandler(func() bool { return db.Ping() == nil }),
		Todos:       http.NewTodosHandler(),
	}
	err = http.NewRouter(endpoints).ServeOn(3000)
	if err != nil {
		panic(err)
	}
}

func createPact() dsl.Pact {
	return dsl.Pact{
		Provider: "API",
	}
}
