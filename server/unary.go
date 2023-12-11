package main

import (
	"context"

	pb "github.com/adarsh2858/grpc-usecases/proto"
)

func (c *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "hello",
	}, nil
}
