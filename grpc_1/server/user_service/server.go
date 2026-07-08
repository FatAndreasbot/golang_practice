package userservice

import (
	"context"
	"proto"
	"regexp"
	db "server/db"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

var emailRegexPool sync.Pool = sync.Pool{
	New: func() any {
		compiledRegex, _ := regexp.Compile(`^[\w\-\.]+(\+[\w\-\.]+)?@([\w-]+\.)+[\w-]{2,}$`)
		return compiledRegex
	},
}

type UserServiceServer struct {
	proto.UnimplementedUserAPIServer
	storage db.DBMock
}

func NewServer() *UserServiceServer {
	return &UserServiceServer{
		storage: *db.NewDBMock(),
	}
}

func (s *UserServiceServer) Create(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	regex := emailRegexPool.Get().(*regexp.Regexp)
	if !regex.MatchString(req.GetEmail()) {
		return nil, status.Error(codes.InvalidArgument, "invalid email")
	}

	if req.GetPassword() != req.GetPasswordConfirm() {
		return nil, status.Error(codes.InvalidArgument, "passwords are not valid")
	}

	newUser := db.User{
		Name:      req.GetName(),
		Password:  req.GetPassword(),
		Role:      db.Role(req.GetRole()),
		CreatedAt: time.Now(),
	}

	id := s.storage.Create(&newUser)

	return &proto.CreateUserResponse{
		Id: id,
	}, nil
}

func (s *UserServiceServer) Get(ctx context.Context, req *proto.SingleUserRequest) (*proto.GetUserResponse, error) {
	UserIDToGet := req.GetId()
	userData, err := s.storage.Retrieve(UserIDToGet)
	if err != nil {
		return nil, err
	}

	return &proto.GetUserResponse{
		Id:        userData.ID,
		Name:      userData.Name,
		Email:     userData.Email,
		Role:      proto.Role(userData.Role),
		CreatedAt: timestamppb.New(userData.CreatedAt),
		UpdatedAt: timestamppb.New(userData.UpdatedAt),
	}, nil
}

func (s *UserServiceServer) Update(ctx context.Context, req *proto.UpdateUserRequest) (*emptypb.Empty, error) {
	userData, err := s.storage.Retrieve(req.GetId())
	if err != nil {
		return nil, err
	}
	if req.GetName() != nil {
		userData.Name = req.GetName().GetValue()
	}
	if req.GetEmail() != nil {
		userData.Email = req.GetEmail().GetValue()
	}
	return &emptypb.Empty{}, nil
}

func (s *UserServiceServer) Delete(ctx context.Context, req *proto.SingleUserRequest) (*emptypb.Empty, error) {
	s.storage.Delete(req.GetId())
	return &emptypb.Empty{}, nil
}
