package main

import (
	grpcPostClient "github.com/racibaz/go-arch/cmd/grpc_client/posts"
	"github.com/racibaz/go-arch/internal/modules/post/presentation/grpc/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// todo make addr configurable via env var or config file
var addr string = "0.0.0.0:51186"

func main() {

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	if err != nil {
		log.Fatalf("Couldn't connect to grpc client: %v\n", err)
	}

	defer conn.Close()
	c := proto.NewPostServiceClient(conn)

	grpcPostClient.CreatePost(c)

}
