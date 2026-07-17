package market

import (
	"proto/user_service"
	"spot_instrument_service/data/models"
)

type MarketStore interface{
	GetMarkets() []*models.Market
	FilteredByUserRoles(userRole user_service.UserRole) []*models.Market
}
