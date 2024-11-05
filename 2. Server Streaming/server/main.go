package main

import (
	"context"
	"fmt"
	"log"
	"net"

	userpb "github.com/Its-Maniaco/grpc-microservice/pb/user/v1"
	wearablepb "github.com/Its-Maniaco/grpc-microservice/pb/wearable/v1"
	"google.golang.org/grpc"
)

// Theses are service_grpc.pb implmentations
type userService struct {
	userpb.UnimplementedUserServiceServer
}

func (u *userService) GetUser(_ context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	return &userpb.GetUserResponse{
		User: &userpb.User{
			Uuid:     req.Uuid,
			FullName: "Its-Maniaco",
		},
	}, nil
}

func main() {
	fmt.Println("Starting server")
	// create listner
	lis, err := net.Listen("tcp", "localhost:9879")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//create grpc server
	grpcServer := grpc.NewServer()

	// register the service to the server
	userpb.RegisterUserServiceServer(grpcServer, &userService{})
	wearablepb.RegisterWearableServiceServer(grpcServer, &wearableService{})

	fmt.Println("Waiting for client request")

	// start (server) server
	grpcServer.Serve(lis)
}
