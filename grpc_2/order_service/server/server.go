package server

import (
	"context"
	"errors"
	"order_service/server/models"
	orderServiceDeclaration "proto/order_service"
	spotInstrumentService "proto/spot_instrument_service"
	"sync"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type OrderServer struct {
	orderServiceDeclaration.UnimplementedOrderServiceServer
	spotInstrumentClient *spotInstrumentService.SpotInstrumentServiceClient
	storage              map[int64]*models.Order
	nextID               int64
	lock                 sync.RWMutex
}

func NewOrderServer() (*OrderServer, error) {
	server := OrderServer{
		storage: make(map[int64]*models.Order),
		nextID:  1,
	}
	// TODO get service address from .env
	conn, err := grpc.NewClient(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.Join(err, errors.New("could not create client"))
	}
	client := spotInstrumentService.NewSpotInstrumentServiceClient(conn)
	server.spotInstrumentClient = &client

	return &server, nil
}

func (s *OrderServer) CreateOrder(ctx context.Context, req *orderServiceDeclaration.CreateOrderRequest) (*orderServiceDeclaration.CreateOrderResponse, error) {
	if len(req.MarketUuid) != 16 {
		return nil, status.Error(codes.InvalidArgument, "the length of the uuid is not 16")
	}

	newOrder := models.Order{
		UserID:     req.GetUserId(),
		MarketUUID: uuid.UUID(req.MarketUuid),
		OrderType:  req.GetOrderType(),
		Price:      int(req.GetPrice()),
		Quantity:   uint(req.GetPrice()),
		Status:     models.CREATED,
	}
	thisOrderID := s.nextID

	s.lock.Lock()
	s.nextID++
	s.storage[thisOrderID] = &newOrder
	s.lock.Unlock()

	go func() {
		time.Sleep(time.Second * 5)

		s.lock.Lock()
		defer s.lock.Unlock()
		if s.storage[thisOrderID].Status == models.CREATED {
			s.storage[thisOrderID].Status = models.PROCESSING
		}
	}()

	return &orderServiceDeclaration.CreateOrderResponse{
		OrderId: thisOrderID,
		Status:  orderServiceDeclaration.ORDER_STATUS_CREATED,
	}, nil
}

func (s *OrderServer) GetOrderStatus(ctx context.Context, req *orderServiceDeclaration.OrderStatusRequest) (*orderServiceDeclaration.OrderStatusResponse, error) {
	order, ok := s.storage[req.GetOrderId()]
	if !ok {
		return nil, status.Error(codes.NotFound, "order was not found")
	}
	if order.UserID != req.GetUserId() {
		return nil, status.Error(codes.PermissionDenied, "this order was created by a different user")
	}

	// in real code i should not ignore this error, ofc
	statusName, _ := order.Status.ToString()

	return &orderServiceDeclaration.OrderStatusResponse{
		Status: orderServiceDeclaration.ORDER_STATUS(orderServiceDeclaration.ORDER_STATUS_value[statusName]),
	}, nil
}
