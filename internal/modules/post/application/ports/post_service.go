package ports

import (
	"context"

	dto "github.com/racibaz/go-arch/internal/modules/post/application/dtos"
	"github.com/racibaz/go-arch/internal/modules/post/domain"
)

// PostService todo it should be separated into command and query services
type PostService interface {
	CreatePost(ctx context.Context, postDto dto.CreatePostInput) error
	GetById(ctx context.Context, id string) (*domain.Post, error)
	// Remove(postID string, userID string) (*domain.Post, error)
}
