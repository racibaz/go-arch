package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/racibaz/go-arch/internal/providers/routes"
	"github.com/racibaz/go-arch/pkg/config"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		<-sigCh
		cancel()
	}()

	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Printf("Listening at %s\n", s.addr)

	grpcServer := grpc.NewServer()

	routes.RegisterGrpcRoutes(grpcServer)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("failed to serve: %v", err)
			cancel()
		}
	}()

	// wait for the shutdown signal
	<-ctx.Done()
	log.Println("Shutting down the server...")
	grpcServer.GracefulStop()
}

func Serve() {
	// Initialize the gRPC server with the address from the configuration
	configs := config.Get()

	grpcServer := NewGRPCServer(fmt.Sprintf("%s:%s", configs.Grpc.Host, configs.Grpc.Port))
	grpcServer.Run()
}
