package main

import (
	"backend/controllers"
	"backend/infrastructures"
	"backend/repositories"
	"backend/services"
	couchbase "github.com/couchbase/gocb/v2"
	"github.com/labstack/gommon/log"
	"sync"
	"time"
)

type IServiceContainer interface {
	InjectPlayerController() controllers.TodoController
}

type kernel struct{}

func NewCouchbaseClient() *couchbase.Cluster {
	cbClient, err := couchbase.Connect(
		"localhost:8091",
		couchbase.ClusterOptions{
			Username: "Admin",
			Password: "123456",
		})
	if err != nil {
		log.Errorf("Error while connecting to cb {}", err)
		panic(err)
	}

	bucket := cbClient.Bucket("Todo")
	err = bucket.WaitUntilReady(3*time.Second, nil)

	return cbClient
}

func (k *kernel) InjectPlayerController() controllers.TodoController {

	couchbaseClient := NewCouchbaseClient()
	couchbaseHandler := &infrastructures.CouchbaseRepository{}
	couchbaseHandler.CbClient = couchbaseClient

	todoRepository := &repositories.TodoRepository{couchbaseHandler}
	todoService := &services.TodoService{todoRepository}
	TodoController := controllers.TodoController{todoService}

	return TodoController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}