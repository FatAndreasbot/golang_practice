package main

import (
	"net"
	"order_service/service_init"
)


func main() {
	listener, err := net.Listen("tcp",":8081")
	if err != nil {
		panic(err)
	}
	grpcServer, err := service_init.GetOrderGRPCServer()
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}

}
