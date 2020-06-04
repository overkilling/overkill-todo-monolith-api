package todo

// Todo represents a single todo task.
type Todo struct {
	Todo string `json:"todo"`
}

// TodosRepository specifies repository functions for getting,
// creating and updating todos.
type TodosRepository interface {
	GetAll() ([]Todo, error)
}
