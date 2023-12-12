package main

import (
	"context"
	"io"
	"log"

	pb "github.com/adarsh2858/grpc-usecases/proto"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Println("Starting Stream")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("couldn't send names; \n %v", err)
	}

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			// log.Print(err.Error())
			log.Printf("error while streaming: \n %s", err)
		}

		log.Printf("Response: %v", response.Message)
	}
	log.Println("Ending Stream")
}
