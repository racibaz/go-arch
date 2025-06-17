package domain

import (
	"errors"
	postValueObject "github.com/racibaz/go-arch/internal/modules/post/domain/value_objects"
	"time"
)

type Post struct {
	ID          string
	Title       string
	Description string
	Status      postValueObject.PostStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewPost(id, title, description string, status postValueObject.PostStatus) (*Post, error) {

	if id == "" {
		return nil, errors.New("id cannot be empty")
	}

	if len(title) < 3 {
		return nil, errors.New("title must be at least 3 characters long")
	}

	if len(description) < 10 {
		return nil, errors.New("title must be at least 3 characters long")
	}

	if !postValueObject.IsValidPostStatus(status) {
		return nil, errors.New("status is not valid")
	}

	return &Post{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}
