package main

import (
	"context"
	"fmt"
	"log"

	api "grpc-client/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := api.NewGreeterClient(conn)

	req := &api.HelloRequest{
		Name: "John",
	}

	resp, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}

	fmt.Println("Server Response:", resp.GetGreeting())
}
