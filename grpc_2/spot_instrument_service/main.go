package main

import (
	"fmt"
	"net"
	spot_instrument "proto/spot_instrument_service"
	"spot_instrument_service/data/store/market"
	"spot_instrument_service/services"
	"spot_instrument_service/services/interceptors"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(auth.UnaryServerInterceptor(
			interceptors.Authenticate,
		)),
	)
	store := market.NewMemMarketStore()

	service := services.NewSpotInstrumentServer(store)
	spot_instrument.RegisterSpotInstrumentServiceServer(
		grpcServer,
		service,
	)

	fmt.Println("listening on localhost:8082")

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}

}
