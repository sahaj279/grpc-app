package main

import (
	"context"
	"fmt"
	"log"
	"net"

	api "grpc-app/api" // Import the generated code

	"google.golang.org/grpc"
)

type server struct {
	api.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *api.HelloRequest) (*api.HelloResponse, error) {
	name := req.GetName()
	response := &api.HelloResponse{
		Greeting: fmt.Sprintf("Hello, %s!", name),
	}
	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterGreeterServer(s, &server{})

	fmt.Println("Server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
