package http_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/lib/pq"
	todoHttp "github.com/overkilling/overkill-todo-monolith-api/http"
	"github.com/overkilling/overkill-todo-monolith-api/postgres"
	"github.com/overkilling/overkill-todo-monolith-api/testcontainers"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationRouter(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests.")
	}

	ctx := context.Background()
	container, err := testcontainers.NewPostgresContainer(ctx)
	if err != nil {
		panic(err)
	}
	defer container.Terminate(ctx)

	ip, err := container.Host(ctx)
	if err != nil {
		panic(err)
	}
	port, err := container.MappedPort(ctx, "5432")
	if err != nil {
		panic(err)
	}

	db, err := postgres.NewDb(
		postgres.DbName("todo"),
		postgres.Credentials("postgres", "postgres"),
		postgres.HostAndPort(ip, port.Int()),
		postgres.SslDisabled())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = postgres.MigrateDB(db, "file://../postgres/migrations")
	if err != nil {
		panic(err)
	}

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:3000/health", nil)
	todoHttp.Router(db).ServeHTTP(res, req)

	content, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, "{\"status\":\"ok\"}", string(content))

	req, _ = http.NewRequest("GET", "http://localhost:3000/todos", nil)
	todoHttp.Router(db).ServeHTTP(res, req)

	content, _ = ioutil.ReadAll(res.Body)
	assert.Equal(t, "[{\"todo\":\"Some task\"}]", string(content))
}
