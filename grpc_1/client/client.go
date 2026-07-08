package main

import (
	"context"
	"fmt"
	"log"
	"proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	conn, err := grpc.NewClient(
		"localhost:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	userAPIClient := proto.NewUserAPIClient(conn)
	chatAPIClient := proto.NewChatAPIClient(conn)

	ctx := context.Background()

	resp, err := userAPIClient.Create(ctx, &proto.CreateUserRequest{
		Name:            "james",
		Email:           "james@mail.com",
		Password:        "james123",
		PasswordConfirm: "james1234",
		Role:            proto.Role_USER,
	})
	if err == nil {
		log.Fatal("i expected here an error")
	} else {
		fmt.Println(err.Error())
	}

	resp, err = userAPIClient.Create(ctx, &proto.CreateUserRequest{
		Name:            "james",
		Email:           "james@mail.com",
		Password:        "james123",
		PasswordConfirm: "james123",
		Role:            proto.Role_USER,
	})

	user, err := userAPIClient.Get(ctx, &proto.SingleUserRequest{Id: resp.GetId()})

	chatAPIClient.SendMessage(ctx, &proto.SendMessageRequest{
		From:      "andy",
		Text:      "grpc is actually not that hard...",
		Timestamp: timestamppb.New(time.Now()),
	})

	fmt.Println(user)
}
