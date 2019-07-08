package cmd

import (
	"context"
	"flag"
	"fmt"

	"github.com/anilkd/go-grpc/pkg/server/grpc"
	v1 "github.com/anilkd/go-grpc/pkg/service/v1"
)

// Config is configuration
type Config struct {
	// the TCP port gRPC server listens to
	GrpcPort string
}

// RunServer runs gRPC server
func RunServer() error {
	ctx := context.Background()

	// get configuration from command-line flag
	var cfg Config
	flag.StringVar(&cfg.GrpcPort, "grpc-port", "9090", "gRPC port to bind")
	flag.Parse()

	if len(cfg.GrpcPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: ‘%s’", cfg.GrpcPort)
	}

	v1Api := v1.NewGreeterServer()
	return grpc.RunServer(ctx, v1Api, cfg.GrpcPort)
}
