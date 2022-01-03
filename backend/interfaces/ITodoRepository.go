package interfaces
import (
	"backend/models"
)

type ITodoRepository interface {
	AddTodoRepo(task string) (error)
	GetAllTodosRepo() ([]models.Todo, error)
	DeleteAllTodosRepo() (error)
}