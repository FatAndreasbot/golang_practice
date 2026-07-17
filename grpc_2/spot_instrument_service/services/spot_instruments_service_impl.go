package services

import (
	"context"
	service "proto/spot_instrument_service"
	"proto/user_service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SpotInstrumentServer struct {
	service.UnimplementedSpotInstrumentServiceServer
	user_service_client user_service.UserServiceClient
}

func (s *SpotInstrumentServer) ViewMarkets(ctx context.Context, req *service.ViewMarketsRequest) (*service.ViewMarketsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method ViewMarkets not implemented")
}
