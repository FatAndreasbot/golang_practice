package services

import (
	"context"
	"errors"
	"proto/user_service"
	jwt "user_service/common/utils"
	"user_service/data/models"
	data "user_service/data/stores/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	user_service.UnimplementedUserServiceServer
	store data.UserStore
}

func NewUserService(store data.UserStore) *UserService{
	return &UserService{
		store: store,
	}
}

func (s *UserService) GetUserRole(ctx context.Context, req *emptypb.Empty) (*user_service.UserRoleResponse, error){
	userdata, ok := ctx.Value("userdata").(models.User)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "could not find userID")
	}
	userUUID := userdata.ID

	user, err := s.store.GetByID(userUUID)
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

	jwtToken, err := jwt.EncodeJWT(user)
	if err != nil {
		return nil, status.Error(codes.Internal, "could not generate jwt token")
	}

	return &user_service.LogInResponse{
		JwtToken: jwtToken,
	}, nil
}
