package service_init

import (
	"fmt"
	order_store "order_service/data/store/orders"
	"order_service/services"
	"order_service/services/interceptors/clientside"
	"order_service/services/interceptors/serverside"
	"proto/order_service"
	spot_instrument "proto/spot_instrument_service"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const SPOT_INSTRUMENT_SERVICE_ADDRESS string = "localhost:8082"

func GetOrderGRPCServer() (*grpc.Server, error){
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(auth.UnaryServerInterceptor(serverside.Authenticate)),
	)
	spotInstrumentConn, err := grpc.NewClient(
		SPOT_INSTRUMENT_SERVICE_ADDRESS,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(clientside.AuthInterceptor),
	)
	if err != nil {
		return nil, err
	}
	spotInstrumentClient := spot_instrument.NewSpotInstrumentServiceClient(spotInstrumentConn)

	service := services.NewOrderService(
		order_store.NewOrderMemStore(),
		spotInstrumentClient,
	)

	order_service.RegisterOrderServiceServer(
		grpcServer,
		service,
	)

	fmt.Println("listening on localhost:8081")

	return grpcServer, nil
}
