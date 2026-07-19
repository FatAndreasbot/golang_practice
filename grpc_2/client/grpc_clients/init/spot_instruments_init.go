package grpcInit

import (
	"client/grpc_clients/interceptors"
	spot_instrument "proto/spot_instrument_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// tbh this could be a generic function...
func InitSpotInstrumentServiceClient(address string) (spot_instrument.SpotInstrumentServiceClient, error){
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(interceptors.AuthInterceptor()),
	)
	if err != nil {
		return nil, err
	}

	client := spot_instrument.NewSpotInstrumentServiceClient(conn)
	return client, nil
}
