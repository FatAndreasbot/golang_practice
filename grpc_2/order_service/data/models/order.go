package models

import (
	"proto/order_service"

	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/type/money"
)

type Order struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	MarketUUID uuid.UUID
	OrderType  string
	Price      *money.Money
	Quantity   float64
	Status     order_service.OrderStatus
}


func NewOrder(owner *User, marketID uuid.UUID, orderType string, price *money.Money, quantity float64) *Order {
	return &Order{
		ID: uuid.New(),
		UserID: owner.ID,
		MarketUUID: marketID,
		OrderType: orderType,
		Price: price,
		Quantity: quantity,
		Status: order_service.OrderStatus_ORDER_STATUS_CREATED,
	}
}
