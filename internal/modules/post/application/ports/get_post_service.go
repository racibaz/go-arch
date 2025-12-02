package ports

import (
	"context"
	"github.com/racibaz/go-arch/internal/modules/post/domain"
)

// GetPostService
type GetPostService interface {
	GetPostByID(ctx context.Context, id string) (*domain.Post, error)
}
