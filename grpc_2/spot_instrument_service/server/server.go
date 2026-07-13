package server

import (
	"context"
	service "proto/spot_instrument_service"
	"spot_instrument_service/server/models"
	"sync"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SpotInstrumentServer struct {
	service.UnimplementedSpotInstrumentServiceServer
	store  map[uuid.UUID]*models.Market
	lock   sync.RWMutex
	doOnce sync.Once
}

func (s *SpotInstrumentServer) ViewMarkets(ctx context.Context, req *service.ViewMarketsRequest) (*service.ViewMarketsResponse, error) {
	protoRole := req.GetUserRoles()
	role, err := models.GetRoleFromName(
		service.UserRole_name[int32(protoRole)],
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "given role is not supported\n%v", err)
	}

	allowedMarkets := make([]*service.ViewMarketsResponse_Market, 0)
	s.lock.RLock()

	for marketUUID, market := range s.store {
		if !market.IsValid() {
			continue
		}

		_, givenRoleIsAllowed := market.AllowedRoles[role]
		if givenRoleIsAllowed {
			allowedMarkets = append(allowedMarkets, &service.ViewMarketsResponse_Market{
				MarketUuid: marketUUID[:],
				MarketName: market.Name,
			})
		}
	}

	s.lock.RUnlock()
	return &service.ViewMarketsResponse{
		Markets: allowedMarkets,
	}, nil
}

func (s *SpotInstrumentServer) AddMarket(newMarket *models.Market) {
	newUUID := uuid.New()
	s.lock.Lock()
	defer s.lock.Unlock()
	s.store[newUUID] = newMarket
}

func NewSpotInstrumentServer() *SpotInstrumentServer {
	newServer := SpotInstrumentServer{
		store: make(map[uuid.UUID]*models.Market, 0),
	}
	// TODO fill the server with mock data

	newServer.doOnce.Do(func() {
		uuid.EnableRandPool()
	})

	return &newServer
}
