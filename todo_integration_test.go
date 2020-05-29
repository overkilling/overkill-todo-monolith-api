package todo

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/overkilling/overkill-todo-monolith-api/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
)

func TestIntegrationRouter(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests.")
	}

	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "postgres",
			"POSTGRES_PASSWORD": "postgres",
			"POSTGRES_DB":       "todo",
		},
	}
	postgresC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		panic(err)
	}
	defer postgresC.Terminate(ctx)

	ip, err := postgresC.Host(ctx)
	if err != nil {
		panic(err)
	}
	port, err := postgresC.MappedPort(ctx, "5432")
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

	// TODO: remove sleep waiting for db to be ready
	time.Sleep(5 * time.Second)

	res := httptest.NewRecorder()
	httpReq, _ := http.NewRequest("GET", "http://localhost:3000/health", nil)

	Router(db).ServeHTTP(res, httpReq)

	content, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, "{\"status\":\"ok\"}", string(content))
}
