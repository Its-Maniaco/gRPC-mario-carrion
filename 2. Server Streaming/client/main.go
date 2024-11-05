package main

import (
	"context"
	"fmt"
	"io"
	"log"

	wearablepb "github.com/Its-Maniaco/grpc-microservice/pb/wearable/v1"
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

	client := wearablepb.NewWearableServiceClient(conn)

	stream, err := client.BeatsPerMinute(context.Background(), &wearablepb.BeatsPerMinuteRequest{
		Uuid: "Its-Maniaco",
	})
	if err != nil {
		log.Fatalln("client error: ", err)
	}

	{
		for {
			strmBalue, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatalln("stream recv: ", err)
			}
			fmt.Println(strmBalue.Value)
		}
	}
}
