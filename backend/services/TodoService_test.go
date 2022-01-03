package services

import (
	"backend/interfaces/mock"
	"backend/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTodoService_AddTodo(t *testing.T) {

	//Given
	todoRepository := new(mock.ITodoRepository)
	todoRepository.On("AddTodoRepo", "test-todo").Return(nil)
	todoService := TodoService{todoRepository}
	expectedResult := error(nil)

	//When
	actualResult := todoService.AddTodo("test-todo")

	//Result
	assert.Equal(t, expectedResult, actualResult)
}

func TestTodoService_GetAllTodos(t *testing.T) {

	//Given
	todoList := []models.Todo{models.Todo{Task: "test-task", Id: "1"}, models.Todo{Task: "test-task2", Id: "2"}}
	expectedErr := error(nil)
	todoRepository := new(mock.ITodoRepository)
	todoRepository.On("GetAllTodosRepo", ).Return(todoList, nil)
	todoService := TodoService{todoRepository}
	expectedResult := []string{todoList[0].Task, todoList[1].Task}


	//When
	actualResult, actualErr := todoService.GetAllTodos()

	//Result
	assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, expectedErr, actualErr)
}

func TestTodoService_DeleteAllTodos(t *testing.T) {

	//Given
	todoRepository := new(mock.ITodoRepository)
	todoRepository.On("DeleteAllTodosRepo", ).Return(nil)
	todoService := TodoService{todoRepository}
	expectedResult := error(nil)

	//When
	actualResult := todoService.DeleteAllTodos()

	//Result
	assert.Equal(t, expectedResult, actualResult)
}