package services

import (
	"context"
	"order_service/data/models"
	store_orders "order_service/data/store/orders"
	"proto/order_service"
	spot_instrument "proto/spot_instrument_service"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderService struct {
	order_service.UnimplementedOrderServiceServer
	spotInstrumentClient spot_instrument.SpotInstrumentServiceClient
	orderStorage         store_orders.OrderStore
}

func NewOrderService(
	storage store_orders.OrderStore,
	spotInstrumentServiceClient spot_instrument.SpotInstrumentServiceClient,
) *OrderService {
	return &OrderService{
		orderStorage: storage,
		spotInstrumentClient: spotInstrumentServiceClient,
	}
}


func (os *OrderService) OrderStatus(
	ctx context.Context, req *order_service.OrderStatusRequest,
) (*order_service.OrderStatusResponse, error) {
	userdata, ok := ctx.Value("userdata").(models.User)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "could not find userID")
	}
	orderUUID, err := uuid.Parse(req.GetOrderUuid())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "could not parse UUID")
	}
	order, err := os.orderStorage.GetByID(orderUUID)
	if err != nil {
		return nil, status.Error(codes.NotFound, "could not find order")
	}

	if order.UserID != userdata.ID {
		return nil, status.Error(codes.NotFound, "could not find order")
	}

	return &order_service.OrderStatusResponse{
		Status: order.Status,
	}, nil
}

func (os *OrderService) CreateOrder(
	ctx context.Context, req *order_service.CreateOrderRequest,
) (*order_service.CreateOrderResponse, error) {

	userdata, ok := ctx.Value("userdata").(models.User)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "could not find userID")
	}

	availableMarkets, err := os.spotInstrumentClient.ViewMarkets(ctx, &spot_instrument.ViewMarketsRequest{
		UserRole: userdata.Role,
	})
	if err != nil {
		return nil, status.Error(codes.Unavailable, err.Error())
	}

	marketUUID, err := uuid.Parse(req.GetMarketUuid())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "market uuid is not correct")
	}

	hasRight := false
	for _, markerResponseData := range availableMarkets.GetMarkets() {
		if markerResponseData.MarketUuid == req.GetMarketUuid(){
			hasRight = true
			break
		}
	}
	if !hasRight {
		return nil, status.Error(codes.PermissionDenied, "requested market is unavailable")
	}

	newOrder := models.NewOrder(
		&userdata,
		marketUUID,
		req.GetOrderType(),
		req.GetPrice(),
		req.GetQuantity(),
	)

	orderUUID, err := os.orderStorage.AddOrder(newOrder)
	if err != nil {
		return nil, status.Error(codes.Internal, "could not save order")
	}

	return &order_service.CreateOrderResponse{
		OrderId: orderUUID.String(),
		Status:  newOrder.Status,
	}, nil
}
