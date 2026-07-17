package main

import (
	"fmt"
	"net"
	"proto/user_service"
	"user_service/common/mocks"
	data "user_service/data/stores/user"
	"user_service/services"
	"user_service/services/interceptors"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"google.golang.org/grpc"
)

func main(){
	listener, err := net.Listen("tcp", ":8080")
	if err != nil{
		panic(err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(auth.UnaryServerInterceptor(interceptors.Authenticate)),
	)
	store := data.NewUserMemStore()
	mocks.FillMockUsers(store)

	service := services.NewUserService(store)
	user_service.RegisterUserServiceServer(grpcServer, service)

	fmt.Println("listening on localhost:8080")

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
