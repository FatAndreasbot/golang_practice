package server

import (
	"context"
	"errors"
	"order_service/server/models"
	orderServiceDeclaration "proto/order_service"
	spotInstrumentService "proto/spot_instrument_service"
	"t1"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type OrderServer struct {
	orderServiceDeclaration.UnimplementedOrderServiceServer
	spotInstrumentClient spotInstrumentService.SpotInstrumentServiceClient
	orderStorage         *t1.Cache[int64, *models.Order]
	userStorage          *t1.Cache[int64, *models.User]

	nextID int64
}

func NewOrderServer() (*OrderServer, error) {
	server := OrderServer{
		orderStorage: t1.NewCache[int64, *models.Order](),
		userStorage:  t1.NewCache[int64, *models.User](),
		nextID:       1,
	}
	// TODO get service address from .env
	conn, err := grpc.NewClient(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.Join(err, errors.New("could not create client"))
	}
	client := spotInstrumentService.NewSpotInstrumentServiceClient(conn)
	server.spotInstrumentClient = client

	// TODO fill userStorage with mock data

	return &server, nil
}

func (s *OrderServer) CreateOrder(ctx context.Context, req *orderServiceDeclaration.CreateOrderRequest) (*orderServiceDeclaration.CreateOrderResponse, error) {
	if len(req.MarketUuid) != 16 {
		return nil, status.Error(codes.InvalidArgument, "the length of the uuid is not 16")
	}

	user, ok := s.userStorage.Get(req.GetUserId())
	if !ok {
		return nil, status.Error(codes.NotFound, "user was not found")
	}

	s.spotInstrumentClient.ViewMarkets(ctx, &spotInstrumentService.ViewMarketsRequest{
		UserRoles: user.Role,
	})

	newOrder := models.Order{
		UserID:     req.GetUserId(),
		MarketUUID: uuid.UUID(req.MarketUuid),
		OrderType:  req.GetOrderType(),
		Price:      int(req.GetPrice()),
		Quantity:   uint(req.GetPrice()),
		Status:     models.CREATED,
	}
	thisOrderID := s.nextID

	s.spotInstrumentClient.ViewMarkets(ctx, &spotInstrumentService.ViewMarketsRequest{})

	s.nextID++
	s.orderStorage.Set(thisOrderID, &newOrder)

	go func() {
		time.Sleep(time.Second * 5)

		order, _ := s.orderStorage.Get(thisOrderID)
		if order.Status == models.CREATED {
			order.Status = models.PROCESSING
		}
	}()

	return &orderServiceDeclaration.CreateOrderResponse{
		OrderId: thisOrderID,
		Status:  orderServiceDeclaration.ORDER_STATUS_CREATED,
	}, nil
}

func (s *OrderServer) GetOrderStatus(ctx context.Context, req *orderServiceDeclaration.OrderStatusRequest) (*orderServiceDeclaration.OrderStatusResponse, error) {
	order, ok := s.orderStorage.Get(req.GetOrderId())
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
