package services

import (
	"backend/interfaces"
)
type TodoService struct {
	interfaces.ITodoRepository
}

func (service *TodoService) AddTodo(task string) (error) {
	err := service.AddTodoRepo(task)
	return err
}

func (service *TodoService) GetAllTodos() ([]string, error) {
	todoList, err := service.GetAllTodosRepo()

	var taskArray = make([]string, len(todoList))
	for i := 0; i < len(todoList); i++ {
		taskArray[i] = todoList[i].Task
	}
	return taskArray, err
}

func (service *TodoService) DeleteAllTodos() (error) {
	err := service.DeleteAllTodosRepo()
	return err
}