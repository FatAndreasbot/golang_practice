package handlers

import (
	grpcInit "client/grpc_clients/init"
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"proto/order_service"
	"sync"

	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/type/money"
)

var initOrderServiceOnce sync.Once
var globalOrderServicehandler orderServiceHandler

const ORDER_SERVICE_ADDRESS = "localhost:8081"

type orderServiceHandler struct {
	grpcClient order_service.OrderServiceClient
	orderStore map[int]uuid.UUID
	mu         sync.RWMutex
	nextID     int
}

func (h *orderServiceHandler) handleList() string {
	var result strings.Builder
	h.mu.RLock()
	defer h.mu.RUnlock()

	for orderNumber, orderUUID := range h.orderStore {
		fmt.Fprintf(&result, "%d - %q", orderNumber, orderUUID.String())
	}
	return result.String()
}

func (h *orderServiceHandler) handleStatus(stringID string) (order_service.OrderStatus, error) {
	ctx := context.Background()
	var result order_service.OrderStatus
	notFoundError := errors.New("given order id was not found")
	var orderUUID uuid.UUID
	intID, err := strconv.Atoi(stringID)
	if err == nil {
		h.mu.RLock()
		var ok bool
		orderUUID, ok = h.orderStore[intID]
		h.mu.RUnlock()
		if !ok {
			return result, notFoundError
		}
	} else {
		orderUUID, err = uuid.Parse(stringID)
		if err != nil {
			return result, errors.New("could not parse given UUID. Try to give a simple ID")
		}
	}
	statusResp, err := h.grpcClient.OrderStatus(ctx, &order_service.OrderStatusRequest{
		OrderUuid: orderUUID.String(),
	})
	if err != nil {
		return result, err
	}
	return statusResp.GetStatus(), nil
}

func (h *orderServiceHandler) handleCreate(marketUUID string, orderType string, price string, quantity string) (int, uuid.UUID, error) {
	ctx := context.Background()
	var orderID int
	var orderUUID uuid.UUID

	realQuantiry, err := strconv.ParseFloat(quantity, 64)
	if err != nil {
		return orderID, orderUUID, errors.New("could not parse quanitiy")
	}

	priceSplit := strings.Split(price, ".")
	var priceUnits, priceNanos int64
	if 1 > len(priceSplit) && len(priceSplit) > 2 {
		return orderID, orderUUID, errors.New("could not parse price")
	}
	priceUnits, err = strconv.ParseInt(priceSplit[0], 10, 64)
	if err != nil {
		return orderID, orderUUID, errors.New("could not parse price")
	}
	if len(priceSplit) == 2 {
		priceNanos, err = strconv.ParseInt(priceSplit[1], 10, 32)
		if err != nil {
			return orderID, orderUUID, errors.New("could not parse price")
		}
	}

	realPrice := money.Money{
		CurrencyCode: "RUB",
		Units:        priceUnits,
		Nanos:        int32(priceNanos),
	}

	createResp, err := h.grpcClient.CreateOrder(ctx, &order_service.CreateOrderRequest{
		MarketUuid: marketUUID,
		OrderType:  orderType,
		Price:      &realPrice,
		Quantity:   realQuantiry,
	})
	if err != nil {
		return orderID, orderUUID, err
	}

	orderUUID, _ = uuid.Parse(createResp.GetOrderId())
	h.mu.Lock()
	orderID = h.nextID
	h.nextID++
	h.orderStore[orderID] = orderUUID
	h.mu.Unlock()

	return orderID, orderUUID, nil
}

func GetOrderServiceHandler() *orderServiceHandler {
	initOrderServiceOnce.Do(func() {
		client, err := grpcInit.InitOrderServiceClient(ORDER_SERVICE_ADDRESS)
		if err != nil {
			panic(err)
		}

		globalOrderServicehandler = orderServiceHandler{
			grpcClient: client,
			orderStore: make(map[int]uuid.UUID),
			nextID:     1,
		}
	})
	return &globalOrderServicehandler
}

func AddOrderServiceHandlers(dispatcher *CommandDispatcher) {
	dispatcher.AddCommand("createorder", "to create a new order [marketUUID] [orderType(string)] [price(float)] [quantity(float)]", func(params ...string) (string, error) {
		if len(params) != 4 {
			return "", errors.New("expected 4 params\n[marketUUID] [orderType(string)] [price(float)] [quantity(float)]")
		}
		orderID, orderUUID, err := GetOrderServiceHandler().handleCreate(
			params[0],
			params[1],
			params[2],
			params[3],
		)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%d - %s", orderID, orderUUID.String()), nil
	})
	dispatcher.AddCommand("orderstatus", "'Order ID/UUID' to get the order status", func(params ...string) (string, error) {
		if len(params) != 1 {
			return "", errors.New("expected only 1 param")
		}
		handler := GetOrderServiceHandler()

		status, err := handler.handleStatus(params[0])
		if err != nil {
			return "", err
		}

		return status.String(), nil
	})
	dispatcher.AddCommand("listorders", "to list all order id/UUIDs", func(params ...string) (string, error) {
		handler := GetOrderServiceHandler()
		return handler.handleList(), nil
	})
}
