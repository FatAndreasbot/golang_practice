package main

import (
	"net"
	"proto/user_service"
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
	service := services.NewUserService(data.NewUserMemStore())
	user_service.RegisterUserServiceServer(grpcServer, service)

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
