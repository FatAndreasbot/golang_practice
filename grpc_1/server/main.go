package main

import (
	chatservice "chat_service"
	"log"
	"net"
	"proto"
	userservice "user_service"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterUserAPIServer(grpcServer, userservice.NewServer())
	proto.RegisterChatAPIServer(grpcServer, chatservice.NewServer())

	log.Printf("server is running at %v\n", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("server crashed, %v\n", err)
	}

}
