package grpcInit

import (
	"client/grpc_clients/interceptors"
	"proto/order_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitOrderServiceClient(address string, ) (order_service.OrderServiceClient, error) {
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(interceptors.AuthInterceptor()),
	)
	if err != nil {
		return nil, err
	}

	client := order_service.NewOrderServiceClient(conn)

	return client, nil
}
