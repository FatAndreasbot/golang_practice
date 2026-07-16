package userservice

import (
	grpcInit "client/grpc_clients/init"
	"client/handlers"
	"client/storage"
	"context"
	"errors"
	"proto/user_service"
	"sync"
)

var once sync.Once
var handler userServiceHandler

const USER_SERVICE_ADDRESS = "localhost:8080"

type userServiceHandler struct {
	grpcClient user_service.UserServiceClient
	tokenStore *storage.TokenStore
}

func GetUserServiceHandler() *userServiceHandler{
	once.Do(func(){
		client, tokenStore, err := grpcInit.InitUserServiceClient(USER_SERVICE_ADDRESS)
		if err != nil {
			panic(err)
		}

		handler = userServiceHandler{
			grpcClient: client,
			tokenStore: tokenStore,
		}
	})

	return &handler
}

func (h *userServiceHandler) handleLogin(username, password string) error {
	ctx := context.Background()
	resp, err := h.grpcClient.LogIn(ctx, &user_service.LogInRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		return err
	}
	h.tokenStore.Set(resp.GetJwtToken())
	return nil
}

func AddUserServiceHandlers(dispatcher *handlers.CommandDispatcher){
	dispatcher.AddCommand("login", func(params ...string) (string, error) {
		if len(params) != 2 {
			return "", errors.New("expected 2 parameters. username password")
		}
		handler := GetUserServiceHandler()
		username, password := params[0], params[1]
		err := handler.handleLogin(username, password)
		if err != nil {
			return "error during login", err
		}
		return "successful login", nil
	})
}
