package main

import (
	"log"
	"net"

	pb "github.com/adarsh2858/grpc-usecases/proto"

	"google.golang.org/grpc"
)

const (
	port = ":3000"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Something went wrong;\n %v", err)
	}

	g := grpc.NewServer()
	pb.RegisterGreetServiceServer(g, &helloServer{})
	log.Printf("Server started at %v", l.Addr())

	if err := g.Serve(l); err != nil {
		log.Fatalf("grpc server could not be served;\n %v", err)
	}
}
