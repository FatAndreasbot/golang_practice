package orders

import (
	"order_service/data/models"

	"github.com/google/uuid"
)

type OrderStore interface{
	AddOrder(*models.Order) (uuid.UUID, error)
	GetByID(uuid.UUID) (*models.Order, error)
}
