package main

import (
	"code/service"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	server := grpc.NewServer()                                      // Create a gRPC server
	service.RegisterUserServiceServer(server, &service.UserService) // Register the UserService

	// Listen for incoming gRPC connections on port 8002
	lis, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatal("Failed to listen ", err)
	}

	// Start serving gRPC requests
	err = server.Serve(lis)
	if err != nil {
		log.Fatal("Failed to serve ", err)
	}
	fmt.Println("Started service successfully")
}
