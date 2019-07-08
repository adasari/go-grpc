package v1

import (
	"context"
	"log"

	apiv1 "github.com/anilkd/go-grpc/pkg/api/v1"
)

type server struct{}

// NewGreeterServer creates GreeterServer
func NewGreeterServer() apiv1.GreeterServer {
	return &server{}
}

// SayHello implements apiv1.GreeterServer interface method
func (s *server) SayHello(ctx context.Context, req *apiv1.HelloRequest) (*apiv1.HelloReply, error) {
	log.Printf("Received the request msg: %v", req)
	return &apiv1.HelloReply{Message: "Hello" + req.Message}, nil
}
