package chatservice

import (
	"context"
	"fmt"
	"proto"
	"t1"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ChatServiceServer struct {
	proto.UnimplementedChatAPIServer
	store  *t1.Cache[int64, *[]string]
	lastID int64
}

func NewServer() *ChatServiceServer {
	return &ChatServiceServer{
		store:  t1.NewCache[int64, *[]string](),
		lastID: 1,
	}
}

func (s *ChatServiceServer) Create(ctx context.Context, req *proto.CreateChatRequest) (*proto.CreateChatResponse, error) {
	chatUsernames := req.GetUsernames()
	s.store.Set(s.lastID, &chatUsernames)
	returnValue := proto.CreateChatResponse{
		Id: s.lastID,
	}
	s.lastID++
	return &returnValue, nil
}
func (s *ChatServiceServer) Delete(ctx context.Context, req *proto.DeleteChatRequest) (*emptypb.Empty, error) {
	toDeleteID := req.GetId()
	s.store.Delete(toDeleteID)

	return &emptypb.Empty{}, nil
}

func (s *ChatServiceServer) SendMessage(ctx context.Context, req *proto.SendMessageRequest) (*emptypb.Empty, error) {
	message := req.GetText()
	sender := req.GetFrom()
	fmt.Printf("%s send a message with text: %q", sender, message)

	return &emptypb.Empty{}, nil
}
