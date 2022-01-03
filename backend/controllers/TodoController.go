package controllers

import (
	"backend/interfaces"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type TodoController struct {
	interfaces.ITodoService
}

type TodoDTO struct {
	Task string `json:"task"`
}

func (controller *TodoController) AddTodoToDB(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")

	res.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	b, err := ioutil.ReadAll(req.Body)
	err = req.Body.Close()
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	// Unmarshal
	var msg TodoDTO
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	err = controller.AddTodo(msg.Task)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	message, err := json.Marshal(msg.Task + " was added")
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	_, err = res.Write(message)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	res.WriteHeader(200)
	return

}

func (controller *TodoController) GetAllTodosFromDB(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")

	res.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	todos, err := controller.GetAllTodos()
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	output, err := json.Marshal(todos)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	res.Header().Set("content-type", "application/json")
	_, err = res.Write(output)

	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	return
}

func (controller *TodoController) DeleteAllTodosFromDB(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	res.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
	err := controller.DeleteAllTodos()
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	message, err := json.Marshal("All Todos Deleted")
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	_, err = res.Write(message)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	res.WriteHeader(200)
	return
}
