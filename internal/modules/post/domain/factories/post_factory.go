package factories

import (
	"github.com/racibaz/go-arch/internal/modules/post/domain"
	"time"
)

// PostFactory is a factory for creating Post instances.
type PostFactory struct {
}

// New initializes a new PostFactory and returns a Post with default values.
func New(id, title, description, content string, status domain.PostStatus, createdAt, updatedAt time.Time) (*domain.Post, error) {

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
		return nil, err
	}

	return &post, nil
}
