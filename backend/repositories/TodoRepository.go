package repositories

import (
	"backend/interfaces"
	"backend/models"
	"github.com/google/uuid"
)

type TodoRepository struct {
	interfaces.IDbHandler
}

func (repository *TodoRepository) AddTodoRepo(task string) (error) {

	id := uuid.New().String()

	newTodo := models.Todo{
		Id: id,
		Task: task,
	}

	err := repository.Add(newTodo)
	return err
}


func (repository *TodoRepository) GetAllTodosRepo() ([]models.Todo, error) {
	query := "SELECT *"
	todos, err := repository.Query(query)
	return todos, err
}

func (repository *TodoRepository) DeleteAllTodosRepo() (error) {
	query := "Delete"
	_, err := repository.Query(query)
	return  err
}