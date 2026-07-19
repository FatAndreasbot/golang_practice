package handlers

import (
	grpcInit "client/grpc_clients/init"
	"client/storage"
	"context"
	"errors"
	"fmt"
	"proto/user_service"
	"sync"

	"google.golang.org/protobuf/types/known/emptypb"
)

var createUserServiceSingletonOnce sync.Once
var userServiceSingleton userServiceHandler

const USER_SERVICE_ADDRESS = "localhost:8080"

type userServiceHandler struct {
	grpcClient user_service.UserServiceClient
}

func GetUserServiceHandler() *userServiceHandler{
	createUserServiceSingletonOnce.Do(func(){
		client, err := grpcInit.InitUserServiceClient(USER_SERVICE_ADDRESS)
		if err != nil {
			panic(err)
		}

		userServiceSingleton = userServiceHandler{
			grpcClient: client,
		}
	})

	return &userServiceSingleton
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
	storage.GetTokenStore().Set(resp.GetJwtToken())
	return nil
}

func (h *userServiceHandler) handleGetRole() (string, error) {
	ctx := context.Background()
	resp, err := h.grpcClient.GetUserRole(ctx, &emptypb.Empty{})
	if err != nil {
		return "", err
	}

	return user_service.UserRole_name[int32(resp.GetUserRole())], nil
}

func AddUserServiceHandlers(dispatcher *CommandDispatcher){
	dispatcher.AddCommand("login", "username, password", func(params ...string) (string, error) {
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

	dispatcher.AddCommand("getrole", "returns your user role", func(s ...string) (string, error) {
		handler := GetUserServiceHandler()
		role, err := handler.handleGetRole()
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("your role is %q", role), nil
	})
}
