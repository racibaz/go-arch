package ports

import (
	usecaseInputs "github.com/racibaz/go-arch/internal/modules/post/application/usecases/inputs"
	"github.com/racibaz/go-arch/internal/modules/post/domain"
)

type PostService interface {
	CreatePost(postDto usecaseInputs.CreatePostInput) error //TODO it should get dto
	GetById(id string) (*domain.Post, error)
}
