package ports

import (
	"github.com/racibaz/go-arch/internal/modules/post/application/dtos"
	"github.com/racibaz/go-arch/internal/modules/post/domain"
)

// PostService todo it should be separated into command and query services
type PostService interface {
	CreatePost(postDto dto.CreatePostInput) error
	GetById(id string) (*domain.Post, error)
	//Remove(postID string, userID string) (*domain.Post, error)
}
