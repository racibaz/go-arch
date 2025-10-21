package ports

import (
	"context"
	"github.com/racibaz/go-arch/internal/modules/post/domain"
)

// PostRepository Secondary port: PostRepository interface
type PostRepository interface {
	Save(ctx context.Context, post *domain.Post) error
	GetByID(ctx context.Context, id string) (*domain.Post, error)
	Update(ctx context.Context, post *domain.Post) error
	Delete(ctx context.Context, id string) error
	List() ([]*domain.Post, error)
	IsExists(title, description string) (bool, error)
}
