package market

import (
	"proto/user_service"
	"spot_instrument_service/data/models"
)

type MarketStore interface{
	SaveMarket(*models.Market) error
	GetMarkets() []*models.Market
	FilteredByUserRoles(userRole user_service.UserRole) []*models.Market
}
