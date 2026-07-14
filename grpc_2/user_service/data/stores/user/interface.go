package data

import (
	"user_service/data/models"

	"github.com/google/uuid"
)

type UserStore interface{
	AddUser(*models.User) error
	GetByID(uuid.UUID) (*models.User, error)
	GetByUsername(string) (*models.User, error)
}
