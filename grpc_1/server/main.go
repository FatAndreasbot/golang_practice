package main

import (
	"context"
	"errors"
	"proto"
	db "server/db"
)

type Server struct {
	proto.UnimplementedUserAPIServer
	storage db.DBMock
}

func (s *Server) Create(ctx context.Context, userdata *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	if userdata.Password != userdata.PasswordConfirm {
		return nil, errors.New("passwords are not valid")
	}

	newUser := db.User{
		Name:     userdata.Name,
		Password: userdata.Password,
		Role:     db.Role(userdata.Role),
	}

	s.storage.Create(&newUser)

	return nil, nil
}

func main() {

}
