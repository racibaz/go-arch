package ports

import (
	"github.com/racibaz/go-arch/internal/modules/post/domain"
)

type PostRepository interface {
	Save(post *domain.Post) error
	GetByID(id string) (*domain.Post, error)
	Update(post *domain.Post) error
	Delete(id string) error
	List() ([]*domain.Post, error)
	IsExists(title, description string) (bool, error)
}
