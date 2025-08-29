package ports

import (
	"github.com/racibaz/go-arch/internal/modules/post/application/commands"
	"github.com/racibaz/go-arch/internal/modules/post/domain"
)

type PostService interface {
	CreatePost(postDto commands.CreatePostInput) error
	GetById(id string) (*domain.Post, error)
	//Remove(postID string, userID string) (*domain.Post, error)
}
