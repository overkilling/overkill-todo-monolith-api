package postgres_test

import (
	"context"
	"database/sql"
	"testing"

	todo "github.com/overkilling/overkill-todo-monolith-api"
	"github.com/overkilling/overkill-todo-monolith-api/postgres"
	"github.com/overkilling/overkill-todo-monolith-api/testcontainers"
	"github.com/stretchr/testify/assert"
)

func TestGetAllTodos(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests.")
	}

	db, err := setupPostgresAndDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := postgres.NewTodosRepository(db)

	todos, err := repository.GetAll()
	if err != nil {
		panic(err)
	}

	assert.ElementsMatch(t, []todo.Todo{
		{Todo: "Some todo"},
		{Todo: "Another todo"},
		{Todo: "Yet another todo"},
	}, todos)
}

func setupPostgresAndDb() (*sql.DB, error) {
	ctx := context.Background()
	container, err := testcontainers.NewPostgresContainer(ctx)
	if err != nil {
		return nil, err
	}

	ip, err := container.Host(ctx)
	if err != nil {
		return nil, err
	}
	port, err := container.MappedPort(ctx, "5432")
	if err != nil {
		return nil, err
	}

	db, err := postgres.NewDb(
		postgres.DbName("todo"),
		postgres.Credentials("postgres", "postgres"),
		postgres.HostAndPort(ip, port.Int()),
		postgres.SslDisabled())
	if err != nil {
		return nil, err
	}

	err = postgres.MigrateDB(db, "file://../postgres/migrations")
	if err != nil {
		return nil, err
	}

	return db, nil
}
