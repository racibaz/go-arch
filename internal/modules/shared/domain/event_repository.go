package domain

// EventRepository Secondary port: EventRepository interface
type EventRepository interface {
	Save(post *Event) error
	GetByID(id string) (*Event, error)
	Update(post *Event) error
	Delete(id string) error
	List() ([]*Event, error)
	IsExists(title, description string) (bool, error)
}
