package main

import (
	"github.com/go-chi/chi"
	"sync"
)

type IChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct{}

func (router *router) InitRouter() *chi.Mux {

	todoController := ServiceContainer().InjectPlayerController()

	r := chi.NewRouter()
	r.Get("/todos", todoController.GetAllTodosFromDB)
	r.Delete("/todos", todoController.DeleteAllTodosFromDB)
	r.Post("/todo", todoController.AddTodoToDB)



	return r
}

var (
	m          *router
	routerOnce sync.Once
)

func ChiRouter() IChiRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}