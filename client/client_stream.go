package main

import (
	"context"
	"log"
	"time"

	pb "github.com/adarsh2858/grpc-usecases/proto"
)

func callSayHelloClientStreaming(client pb.GreetServiceClient, namesList *pb.NamesList) {
	log.Println("Client Streaming Starts")
	stream, _ := client.SayHelloClientStreaming(context.Background())

	for _, name := range namesList.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error sending the request \n %v", err)
		}
		log.Printf("Send request with name: %v", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Print(err.Error())
	}
	log.Println(res.Messages)
	log.Println("Client Streaming End")
}
