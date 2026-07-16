package init

import (
	"client/grpc_clients/interceptors"
	"client/storage"
	"proto/user_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func InitUserServiceClient(address string) (user_service.UserServiceClient, *storage.TokenStore, error){
	tokenStore := storage.TokenStore{}
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(interceptors.AuthInterceptor(&tokenStore)),
	)
	if err != nil {
		return nil, nil, err
	}

	client := user_service.NewUserServiceClient(conn)
	return client, &tokenStore, nil
}
