package main

import (
	"context"
	"log"
	"time"

	pb "github.com/adarsh2858/grpc-usecases/proto"
)

func callSayHello(c pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := c.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Print(err.Error())
		return
	}

	log.Printf("Response: %v", res.Message)
}
