package mock

import (
	"backend/models"
	mock "github.com/stretchr/testify/mock"
)

type ITodoRepository struct {
	mock.Mock
}

func (_m *ITodoRepository) AddTodoRepo(task string) (error) {
	ret := _m.Called(task)

	var r error
	if _, ok := ret.Get(0).(func(string) error); ok {
		r = nil
	} else {
		r = ret.Error(0)
	}

	return r
}

func (_m *ITodoRepository) GetAllTodosRepo() ([]models.Todo, error) {

	var r error
	r = nil
	todos := []models.Todo{models.Todo{Task: "test-task", Id: "1"}, models.Todo{Task: "test-task2", Id: "2"}}
	return todos,r
}

func (_m *ITodoRepository) DeleteAllTodosRepo() (error) {
	err := error(nil)
	return err
}