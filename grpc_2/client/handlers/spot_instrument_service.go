package handlers

import (
	grpcInit "client/grpc_clients/init"
	"context"
	"fmt"
	spot_instrument "proto/spot_instrument_service"
	"proto/user_service"
	"sync"
)

var createSpotInstrumentServicSingletonOnce sync.Once
var spotInstrumentServicSingleton spotInstrumentServiceHandler

const SPOT_INSTRUMENT_SERVICE_ADDRESS = "localhost:8082"


type spotInstrumentServiceHandler struct {
	grpcClient spot_instrument.SpotInstrumentServiceClient
}

func (h *spotInstrumentServiceHandler) handleListMarkets() ([]*spot_instrument.ViewMarketsResponse_Market, error){
	ctx := context.Background()
	userServiceHandler := GetUserServiceHandler()
	roleName, err := userServiceHandler.handleGetRole()
	if err != nil {
		return nil, err
	}
	roleValue := user_service.UserRole_value[roleName]

	listMarketsResponse, err := h.grpcClient.ViewMarkets(ctx, &spot_instrument.ViewMarketsRequest{
		UserRole: user_service.UserRole(roleValue),
	})
	if err != nil {
		return nil, err
	}
	return listMarketsResponse.GetMarkets(), nil
}

func GetSpotInstrumentServiceHandler() *spotInstrumentServiceHandler {
	createSpotInstrumentServicSingletonOnce.Do(func(){
		client, err := grpcInit.InitSpotInstrumentServiceClient(SPOT_INSTRUMENT_SERVICE_ADDRESS)
		if err != nil {
			panic(err)
		}

		spotInstrumentServicSingleton = spotInstrumentServiceHandler{
			grpcClient: client,
		}
	})

	return &spotInstrumentServicSingleton
}

func AddSpotInstrumentServiceHandlers(dispatcher *CommandDispatcher) {
	dispatcher.AddCommand("listmarkets", "not implemented", func(s ...string) (string, error) {
		var result string

		handler := GetSpotInstrumentServiceHandler()
		markets, err := handler.handleListMarkets()
		if err != nil {
			return result, err
		}
		for _, market := range markets{
			result += fmt.Sprintf("%s - %q\n", market.GetMarketName(), market.GetMarketUuid())
		}
		return result, nil
	})
}
