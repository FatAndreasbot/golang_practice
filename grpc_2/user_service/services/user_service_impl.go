package services

import (
	"context"
	"errors"
	"proto/user_service"
	data "user_service/data/stores/user"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	user_service.UnimplementedUserServiceServer
	store data.UserStore
}


func (s *UserService) GetUserRole(ctx context.Context, req *emptypb.Empty) (*user_service.UserRoleResponse, error){
	userIDRaw := ctx.Value("userID")
	if userIDRaw == nil {
		return nil, status.Error(codes.InvalidArgument, "could not find userID")
	}
	userID, ok := userIDRaw.(string)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "could not read userID")
	}
	parsedUserUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, errors.Join(err, errors.New("could not parse into uuid")).Error())
	}
	user, err := s.store.GetByID(parsedUserUUID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, errors.Join(err, errors.New("could not parse into uuid")).Error())
	}

	return &user_service.UserRoleResponse{
		UserRole: user.Role,
	}, nil
}

func (s *UserService) LogIn(ctx context.Context, req *user_service.LogInRequest) (*user_service.LogInResponse, error) {
	user, err := s.store.GetByUsername(req.GetUsername())
	if err != nil {
		return nil, status.Error(
			codes.InvalidArgument, errors.Join(err, errors.New("could not find user by username")).Error(),
		)
	}

	if !user.CheckPassword(req.GetPassword()) {
		return nil, status.Error(
			codes.PermissionDenied, errors.Join(err, errors.New("password incorrect")).Error(),
		)
	}

	return nil, status.Error(codes.Unimplemented, "method LogIn not implemented")
}
