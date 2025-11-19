package posts

import (
	"context"
	"github.com/racibaz/go-arch/internal/modules/post/domain"
	pb "github.com/racibaz/go-arch/internal/modules/post/presentation/grpc/proto"
	"log"
)

const (
	PostAggregate = "posts.Post"
)

func CreatePost(c pb.PostServiceClient) string {

	status, err := domain.PostStatusPublished.ToInt()
	if err != nil {
		log.Fatalf("Invalid status: %v\n", err)
	}

	var post = &pb.CreatePostInput{
		ID:          "7b3a4d03-bcb9-47ce-b721-a156edd406f0",
		UserID:      "7b3a4d03-bcb9-47ce-b721-a156edd406f0",
		Title:       "test title title title grpc",
		Description: "test description description grpc",
		Content:     "test content content content grpc",
		Status:      int32(status),
	}
	res, err := c.CreatePost(context.Background(), post)

	if err != nil {
		log.Fatalf("Could not create post: %v\n", err)
	}

	log.Printf("Post has been created with ID: %s\n", post.ID)

	return res.GetId()
}
