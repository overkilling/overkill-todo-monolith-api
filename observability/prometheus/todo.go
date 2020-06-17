package prometheus

import (
	"time"

	todo "github.com/overkilling/overkill-todo-monolith-api"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var repositoryInstrumentationLabels = []string{"method"}

type prometheusTodoRepo struct {
	instrumentedRepo  todo.TodosRepository
	requestCounter    *prometheus.CounterVec
	errorCounter      *prometheus.CounterVec
	durationHistogram *prometheus.HistogramVec
}

func (repository *prometheusTodoRepo) GetAll() ([]todo.Todo, error) {
	start := time.Now()
	todos, err := repository.instrumentedRepo.GetAll()

	labels := prometheus.Labels{"method": "GetAll"}
	repository.durationHistogram.With(labels).Observe(time.Since(start).Seconds())
	repository.requestCounter.With(labels).Inc()
	if err != nil {
		repository.errorCounter.With(labels).Inc()
	}

	return todos, err
}

// InstrumentTodosRepository instruments a Todos repository with prometheus
// metrics gathering
func InstrumentTodosRepository(repository todo.TodosRepository) todo.TodosRepository {
	requestCounter := promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "todos_repository_requests_total",
			Help: "Total number of requests made to the Todos repository",
		},
		repositoryInstrumentationLabels,
	)
	errorCounter := promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "todos_repository_requests_errors",
			Help: "Total number of errors on requests made to the Todos repository",
		},
		repositoryInstrumentationLabels,
	)
	durationHistogram := promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "todos_repository_requests_duration",
			Help: "Histogram of Todo's repository request durations in seconds",
		},
		repositoryInstrumentationLabels,
	)

	return &prometheusTodoRepo{repository, requestCounter, errorCounter, durationHistogram}
}
