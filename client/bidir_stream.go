package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/adarsh2858/grpc-usecases/proto"
)

func callSayHelloBidirectionalStreaming(client pb.GreetServiceClient, namesStruct *pb.NamesList) {
	log.Print("Bidirectional streaming starts here")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Printf("error while calling bidirectional streaming; \n %v", err)
	}

	// receive the stream of responses from server
	waitch := make(chan struct{})
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("client could not receive data; \n %v", err)
			}

			log.Printf("Response: %v", res)
		}

		close(waitch)
	}()

	// send the names as a stream
	for _, name := range namesStruct.Names {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf(err.Error())
		}

		time.Sleep(2 * time.Second)
	}

	stream.CloseSend()
	<-waitch
	log.Print("Bidirectional streaming ends here")
}
