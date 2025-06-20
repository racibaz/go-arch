package grpc

import (
	"fmt"
	"github.com/racibaz/go-arch/internal/providers/routes"
	"github.com/racibaz/go-arch/pkg/config"
	"google.golang.org/grpc"
	"log"
	"net"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() {
	lis, err := net.Listen("tcp", s.addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Printf("Listening at %s\n", s.addr)

	grpcServer := grpc.NewServer()

	routes.RegisterGrpcRoutes(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}

func Serve() {
	// Initialize the gRPC server with the address from the configuration
	configs := config.Get()

	grpcServer := NewGRPCServer(fmt.Sprintf("%s:%s", configs.Server.GrpcHost, configs.Server.GrpcPort))
	grpcServer.Run()
}
