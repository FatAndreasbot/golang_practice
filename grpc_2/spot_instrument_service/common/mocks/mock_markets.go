package mocks

import (
	"proto/user_service"
	"spot_instrument_service/data/models"
	store "spot_instrument_service/data/store/market"
)

func FillMockMarkets(store store.MarketStore){
	marketOpen := models.NewMarket(
		"open",
		user_service.UserRole_USER_ROLE_ADMIN,
		user_service.UserRole_USER_ROLE_BROKER,
		user_service.UserRole_USER_ROLE_USER,
	)
	marketSpecial := models.NewMarket(
		"special",
		user_service.UserRole_USER_ROLE_BROKER,
	)
	marketTesting := models.NewMarket(
		"test",
		user_service.UserRole_USER_ROLE_ADMIN,
	)

	store.SaveMarket(marketOpen)
	store.SaveMarket(marketSpecial)
	store.SaveMarket(marketTesting)
}
