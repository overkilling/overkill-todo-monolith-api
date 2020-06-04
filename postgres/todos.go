package postgres

import (
	"database/sql"

	todo "github.com/overkilling/overkill-todo-monolith-api"
)

// TodosRepository is a postgres implementation of a
// base todos.TodosRepository
type TodosRepository struct {
	db *sql.DB
}

// NewTodosRepository creates a new Todos repository backed by a
// Postgres db.
func NewTodosRepository(db *sql.DB) *TodosRepository {
	return &TodosRepository{db}
}

// GetAll returns all todos from the database.
func (repository *TodosRepository) GetAll() ([]todo.Todo, error) {
	rows, err := repository.db.Query("SELECT todo FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := make([]todo.Todo, 0)
	for rows.Next() {
		var text string
		if err := rows.Scan(&text); err != nil {
			return nil, err
		}
		todos = append(todos, todo.Todo{Todo: text})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}
