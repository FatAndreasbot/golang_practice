package services

import (
	"context"
	service "proto/spot_instrument_service"
	"spot_instrument_service/data/models"
	"spot_instrument_service/data/store/market"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SpotInstrumentServer struct {
	service.UnimplementedSpotInstrumentServiceServer
	store market.MarketStore
}

func NewSpotInstrumentServer(store market.MarketStore) *SpotInstrumentServer{
	return &SpotInstrumentServer{
		store: store,
	}
}

func (s *SpotInstrumentServer) ViewMarkets(ctx context.Context, req *service.ViewMarketsRequest) (*service.ViewMarketsResponse, error) {
	userdata, ok := ctx.Value("userdata").(models.User)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "could not find userID")
	}
	filterdMarkets := s.store.FilteredByUserRoles(userdata.Role)
	var resultMarktets []*service.ViewMarketsResponse_Market

	for _, market := range filterdMarkets {
		serviceResponseMarket := service.ViewMarketsResponse_Market{
			MarketUuid: market.ID.String(),
			MarketName: market.Name,
		}
		resultMarktets = append(resultMarktets, &serviceResponseMarket)
	}

	return &service.ViewMarketsResponse{
		Markets: resultMarktets,
	}, nil
}
