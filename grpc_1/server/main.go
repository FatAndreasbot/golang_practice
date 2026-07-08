package main

import (
	"context"
	"errors"
	"log"
	"net"
	"proto"
	"regexp"
	db "server/db"
	"sync"
	"time"

	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

var emailRegexPool sync.Pool = sync.Pool{
	New: func() any {
		compiledRegex, _ := regexp.Compile(`^[\w\-\.]+(\+[\w\-\.]+)?@([\w-]+\.)+[\w-]{2,}$`)
		return compiledRegex
	},
}

type Server struct {
	proto.UnimplementedUserAPIServer
	storage db.DBMock
}

func NewServer() *Server {
	return &Server{
		storage: *db.NewDBMock(),
	}
}

func (s *Server) Create(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	regex := emailRegexPool.Get().(*regexp.Regexp)
	if !regex.MatchString(req.GetEmail()) {
		return nil, errors.New("invalid email")
	}

	if req.GetPassword() != req.GetPasswordConfirm() {
		return nil, errors.New("passwords are not valid")
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

func (s *Server) Get(ctx context.Context, req *proto.SingleUserRequest) (*proto.GetUserResponse, error) {
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

func (s *Server) Update(ctx context.Context, req *proto.UpdateUserRequest) (*emptypb.Empty, error) {
	userData, err := s.storage.Retrieve(req.GetId())
	if err != nil {
		return nil, err
	}
	userData.Name = req.GetName().GetValue()
	userData.Email = req.GetEmail().GetValue()

	return &emptypb.Empty{}, nil
}

func (s *Server) Delete(ctx context.Context, req *proto.SingleUserRequest) (*emptypb.Empty, error) {
	s.storage.Delete(req.GetId())
	return &emptypb.Empty{}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterUserAPIServer(grpcServer, NewServer())
	log.Printf("server is running at %v\n", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("server crashed, %v\n", err)
	}
}
