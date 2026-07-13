package models

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type OrderStatus int32

const (
	CREATED OrderStatus = iota
	PROCESSING
	FULFILLED
	CANCELED
)

func (s OrderStatus) ToString() (string, error) {
	switch s {
	case CREATED:
		return "CREATED", nil
	case PROCESSING:
		return "PROCESSING", nil
	case FULFILLED:
		return "FULFILLED", nil
	case CANCELED:
		return "CANCELED", nil
	default:
		return "", errors.New("order value was not found")
	}
}

func GetStatusFromName(name string) (OrderStatus, error) {
	switch name {
	case "CREATED":
		return CREATED, nil
	case "PROCESSING":
		return PROCESSING, nil
	case "FULFILLED":
		return FULFILLED, nil
	case "CANCELED":
		return CANCELED, nil
	default:
		return -1, fmt.Errorf("role %q does not exists", name)
	}
}

type Order struct {
	UserID     int64
	MarketUUID uuid.UUID
	OrderType  string
	Price      int
	Quantity   uint
	Status     OrderStatus
}
