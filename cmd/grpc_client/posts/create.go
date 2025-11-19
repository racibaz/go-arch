package posts

import (
	"context"
	pb "github.com/racibaz/go-arch/internal/modules/post/presentation/grpc/proto"
	"log"
)

func CreatePost(c pb.PostServiceClient) string {

	post := &pb.Post{
		Title:       "test title grpc",
		Description: "test description grpc",
		Content:     "test content grpc",
	}

	res, err := c.CreatePost(context.Background(), post)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Post has been created: %v\n", res)

	return res.Id
}
