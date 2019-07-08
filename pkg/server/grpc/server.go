package grpc

import (
	"context"
	"log"
	"net"

	apiv1 "github.com/anilkd/go-grpc/pkg/api/v1"
	"google.golang.org/grpc"
)

// RunServer runs gRPC server to publish the Greeter service
func RunServer(ctx context.Context, v1Api apiv1.GreeterServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	apiv1.RegisterGreeterServer(server, v1Api)

	// start gRPC server
	log.Println("starting gRPC server at localhost:", port)
	return server.Serve(listen)
}
