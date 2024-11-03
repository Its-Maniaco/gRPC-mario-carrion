package main

import (
	"context"
	"fmt"
	"log"

	userpb "github.com/Its-Maniaco/grpc-microservice/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient("localhost:9879", opts...)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)

	res, err := client.GetUser(context.TODO(), &userpb.GetUserRequest{
		Uuid: "Some string",
	})
	if err != nil {
		log.Fatalf("failed to get user: %v", err)
	}

	fmt.Printf("%+v\n", res)

}
