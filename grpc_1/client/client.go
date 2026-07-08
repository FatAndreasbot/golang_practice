package main

import (
	"context"
	"fmt"
	"log"
	"proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.NewClient("localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := proto.NewUserAPIClient(conn)

	ctx := context.Background()

	resp, err := client.Create(ctx, &proto.CreateUserRequest{
		Name:            "james",
		Email:           "james@mail.com",
		Password:        "james123",
		PasswordConfirm: "james1234",
		Role:            proto.Role_USER,
	})
	if err == nil {
		log.Fatal("i expected here an error")
	}

	user, err := client.Get(ctx, &proto.SingleUserRequest{Id: resp.GetId()})

	fmt.Println(user)
}
