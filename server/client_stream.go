package main

import (
	"io"
	"log"

	pb "github.com/adarsh2858/grpc-usecases/proto"
)

func (h *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if err != nil {
			log.Printf("received request has error: \n %v", err)
			return err
		}

		log.Printf("Got request with name %v", req.Name)
		messages = append(messages, "Hello "+req.Name)
	}
}
