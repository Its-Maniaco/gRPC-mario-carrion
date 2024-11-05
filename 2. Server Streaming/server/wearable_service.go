package main

import (
	"math/rand/v2"
	"time"

	wearablepb "github.com/Its-Maniaco/grpc-microservice/pb/wearable/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type wearableService struct {
	wearablepb.UnimplementedWearableServiceServer
}

func (ws *wearableService) BeatsPerMinute(
	req *wearablepb.BeatsPerMinuteRequest,
	//stream grpc.ServerStreamingServer[wearablepb.BeatsPerMinuteResponse],
	stream wearablepb.WearableService_BeatsPerMinuteServer,
) error {

	for {
		select {
		case <-stream.Context().Done():
			return status.Errorf(codes.Canceled, "Stream Ended")
		default:
			time.Sleep(2 * time.Second)
			value := 30 + rand.Int32N(80)

			err := stream.SendMsg(&wearablepb.BeatsPerMinuteResponse{
				Value:  uint32(value),
				Minute: uint32(time.Now().Second()),
			})
			if err != nil {
				return status.Errorf(codes.Canceled, "Stream Ended")
			}
		}
	}
}
