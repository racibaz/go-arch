package ports

import "github.com/racibaz/go-arch/internal/modules/shared/domain"

// EventRepository Secondary port: EventRepository interface
type EventRepository interface {
	Save(post *domain.Event) error
	GetByID(id string) (*domain.Event, error)
	Update(post *domain.Event) error
	Delete(id string) error
	List() ([]*domain.Event, error)
	IsExists(title, description string) (bool, error)
}
