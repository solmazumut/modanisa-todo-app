package infrastructures

import (
	"backend/interfaces"
	"backend/models"
	couchbase "github.com/couchbase/gocb/v2"
	"github.com/labstack/gommon/log"
)

type CouchbaseRepository struct {
	CbClient *couchbase.Cluster
}


func NewCouchbaseRepositoryAdaptor(cbClient *couchbase.Cluster) interfaces.IDbHandler {
	return &CouchbaseRepository{
		CbClient: cbClient,
	}
}

func (c *CouchbaseRepository) Add(todo models.Todo) (error) {
	_, err := c.CbClient.Bucket("Todo").DefaultCollection().Upsert(todo.Id, todo, &couchbase.UpsertOptions{})

	if err != nil {
		log.Error("An error occurred while inserting document to couchbase ")
		return err
	}

	return nil
}

func (c *CouchbaseRepository) Query(statement string) ([]models.Todo, error) {
	query := statement + " from Todo"
	getResult, err := c.CbClient.Query(query,  &couchbase.QueryOptions{})

	if err != nil {
		log.Error("An error occurred while retrieving documents from couchbase ")
		return nil, err
	}

	if getResult == nil {
		return nil, nil
	}

	var todos []models.Todo
	if getResult != nil {
		for getResult.Next() {
			var row interface{}
			err = getResult.Row(&row)
			if err != nil {
				log.Errorf("Error while getting todo row from couchbase {}", err)
			}


			myMap := row.(map[string]interface{})
			resultMap := myMap["Todo"].(map[string]interface{})
			todo := models.Todo{
				Id: resultMap["Id"].(string),
				Task: resultMap["Task"].(string),
			}
			todos = append(todos,todo)


		}
	}

	return todos, nil
}