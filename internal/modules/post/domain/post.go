package domain

import (
	"errors"
	postValueObject "github.com/racibaz/go-arch/internal/modules/post/domain/value_objects"
	"strings"
	"time"
)

type Post struct {
	ID          string
	Title       string
	Description string
	Content     string
	Status      postValueObject.PostStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// todo it needs to be refactored to use a factory
// todo it needs test
func NewPost(id, title, description, content string, status postValueObject.PostStatus) (*Post, error) {

	// Trim whitespace from the input parameters
	id = strings.TrimSpace(id)
	title = strings.TrimSpace(title)
	description = strings.TrimSpace(description)

	// Validate the input parameters
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}

	if len(title) < 3 {
		return nil, errors.New("title must be at least 3 characters long")
	}

	if len(description) < 10 {
		return nil, errors.New("description must be at least 10 characters long")
	}

	if len(content) < 10 {
		return nil, errors.New("content must be at least 10 characters long")
	}

	if !postValueObject.IsValidPostStatus(status) {
		return nil, errors.New("status is not valid")
	}

	return &Post{
		ID:          id,
		Title:       title,
		Description: description,
		Content:     content,
		Status:      status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}
