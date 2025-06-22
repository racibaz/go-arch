package factories

import (
	"github.com/racibaz/go-arch/internal/modules/post/domain"
	postValueObject "github.com/racibaz/go-arch/internal/modules/post/domain/value_objects"
	"time"
)

// PostFactory is a factory for creating Post instances.
type PostFactory struct {
}

// New initializes a new PostFactory and returns a Post with default values.
func New(id, title, description, content string, status postValueObject.PostStatus, createdAt, updatedAt time.Time) (*domain.Post, error) {

	// This factory method creates a new Post with default values if you want.
	post := domain.Post{
		ID:          id,
		Title:       title,
		Description: description,
		Content:     content,
		Status:      status,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	//validate the post before returning it
	err := post.Validate()
	if err != nil {
		return nil, nil
	}

	return &post, nil
}
