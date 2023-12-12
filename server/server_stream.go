package main

import (
	"log"
	"time"

	pb "github.com/adarsh2858/grpc-usecases/proto"
)

func (h *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("Req from the client: %v", req.Names)

	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}

		err := stream.Send(res)
		if err != nil {
			return err
		}
		time.Sleep(time.Second * 2)
	}

	return nil
}
