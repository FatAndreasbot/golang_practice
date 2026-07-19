package market

import (
	"proto/user_service"
	"spot_instrument_service/data/models"
	"sync"

	"github.com/google/uuid"
)


type MemMarketStore struct {
	store map[uuid.UUID]*models.Market
	mu sync.RWMutex
}

func NewMemMarketStore() *MemMarketStore {
	return &MemMarketStore{
		store: make(map[uuid.UUID]*models.Market),
		mu: sync.RWMutex{},
	}
}

func (s *MemMarketStore) SaveMarket(newMarket *models.Market) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.store[newMarket.ID] = newMarket
	return nil
}

func (s *MemMarketStore) GetMarkets() []*models.Market {
	var result []*models.Market
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, market := range s.store {
		result = append(result, market)
	}

	return result
}

func (s *MemMarketStore) FilteredByUserRoles(
	userRole user_service.UserRole,
) []*models.Market {
	var result []*models.Market
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, market := range s.store {
		_, ok := market.AllowedRoles[userRole]
		if ok{
			result = append(result, market)
		}
	}

	return result
}
