package main

import (
	"context"
	"log"

	v1 "github.com/anilkd/go-grpc/pkg/api/v1"
	"google.golang.org/grpc"
)

func main() {
	// get command-line configuration
	addr := "localhost:9090"
	name := "test"

	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := v1.NewGreeterClient(conn)
	log.Printf("Client connected to the server at address %v, and ready to  make grpc calls.", addr)

	resp, err := client.SayHello(context.Background(), &v1.HelloRequest{Message: name})
	if err != nil {
		log.Fatalf("Failed to make the gRPC call: %v", err)
	}

	log.Printf("Received sayHello response msg: %s", resp.Message)
}
