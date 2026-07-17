package orders

import (
	"errors"
	"math/rand"
	"order_service/data/models"
	"proto/order_service"
	"t1"
	"time"

	"github.com/google/uuid"
)

type OrderMemStore struct{
	store *t1.Cache[uuid.UUID, *models.Order]
}

func NewOrderMemStore() *OrderMemStore{
	return &OrderMemStore{
		store: t1.NewCache[uuid.UUID, *models.Order](),
	}
}

func (s *OrderMemStore) AddOrder(order *models.Order) (uuid.UUID, error) {
	s.store.Set(order.ID, order)

	go func(){
		// simulating order processing
		time.Sleep(time.Second * (time.Duration(rand.Intn(3)+2)))
		order.Status = order_service.OrderStatus_ORDER_STATUS_PROCESSING

		time.Sleep(time.Second * (time.Duration(rand.Intn(8)+2)))
		order.Status = order_service.OrderStatus_ORDER_STATUS_FULFILLED
	}()

	return order.ID, nil
}

func (s *OrderMemStore) GetByID(id uuid.UUID) (*models.Order, error) {
	order, ok := s.store.Get(id)
	if !ok {
		return order, errors.New("order was not found")
	}
	return order, nil
}
