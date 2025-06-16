package ports

import (
	"errors"
	"github.com/racibaz/go-arch/internal/modules/post/domain"
)

var (
	ErrPostNotFound     = errors.New("the post was not found")
	ErrPostAlreadyExist = errors.New("the post already exists")
)

type PostRepository interface {
	Create(post *domain.Post) error
	GetByID(id string) (*domain.Post, error)
	Update(post *domain.Post) error
	Delete(id string) error
	List() ([]*domain.Post, error)
}
