package interfaces

import (
	"backend/models"
)

type IDbHandler interface {
	Add(todo models.Todo) (error)
	Query(statement string) ([]models.Todo, error)
}
