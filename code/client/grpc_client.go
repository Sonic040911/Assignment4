package main

import (
	"code/client/service"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	// Establish a connection to the gRPC server running on port 8002
	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to dial server: ", err)
	}
	defer conn.Close()

	// Create a client
	client := service.NewUserServiceClient(conn)

	// Add User
	addUserResp, err := client.AddUser(context.Background(), &service.User{
		Id:    1,
		Name:  "Yedil",
		Email: "hello123@gmail.com",
	})
	if err != nil {
		log.Fatal("Failed to add user: ", err)
	}
	log.Print("Added user with ID: ", addUserResp.UserId)

	// Get User
	getUserResp, err := client.GetUser(context.Background(), &service.GetUserRequest{
		UserId: 1,
	})
	if err != nil {
		log.Fatal("Failed to get user: ", err)
	}
	log.Print("Retrieved user with ID: ", getUserResp.Id)

	// Check All Users
	listUsersResp, err := client.ListUsers(context.Background(), &service.ListUsersRequest{})
	if err != nil {
		log.Fatal("Failed to list users ", err)
	}
	log.Print("List of users: ", listUsersResp.Users)
}
