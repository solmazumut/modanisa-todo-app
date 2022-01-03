package interfaces

type ITodoService interface {
	AddTodo(task string) (error)
	GetAllTodos() ([]string, error)
	DeleteAllTodos() (error)
}