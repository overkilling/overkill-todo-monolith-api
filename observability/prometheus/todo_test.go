package prometheus_test

import (
	"errors"
	"testing"

	todo "github.com/overkilling/overkill-todo-monolith-api"
	"github.com/overkilling/overkill-todo-monolith-api/observability/prometheus"
	"github.com/stretchr/testify/assert"
)

type testTodoRepo struct {
	shouldFail bool
}

func (repo *testTodoRepo) GetAll() ([]todo.Todo, error) {
	var err error
	if repo.shouldFail {
		err = errors.New("repo failed")
	}
	return []todo.Todo{{Todo: "Test todo"}}, err
}

func TestInstrumentedGetAll(t *testing.T) {
	todoRepository := testTodoRepo{shouldFail: false}
	instrumentedTodoRepository := prometheus.InstrumentTodosRepository(&todoRepository)

	todoRepository.shouldFail = false
	todos, err := instrumentedTodoRepository.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, "Test todo", todos[0].Todo)

	todoRepository.shouldFail = true
	_, err = instrumentedTodoRepository.GetAll()
	assert.Error(t, err)

	content := getPrometheusMetrics(t)
	assert.Contains(t, content, `todos_repository_requests_total{method="GetAll"} 2`)
	assert.Contains(t, content, `todos_repository_requests_errors{method="GetAll"} 1`)
	assert.Contains(t, content, `todos_repository_requests_duration_count{method="GetAll"} 2`)
}
