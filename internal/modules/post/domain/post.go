package domain

import (
	"errors"
	postValueObject "github.com/racibaz/go-arch/internal/modules/post/domain/value_objects"
	"strings"
	"time"
)

var (
	ErrPostNotFound      = errors.New("the post was not found")
	ErrPostAlreadyExists = errors.New("the post already exists")
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

// Validate checks if the Post fields are valid.
func (post *Post) Validate() error {

	// Trim whitespace from the input parameters
	id := strings.TrimSpace(post.ID)
	title := strings.TrimSpace(post.Title)
	description := strings.TrimSpace(post.Description)
	content := strings.TrimSpace(post.Content)

	// Validate the input parameters
	if id == "" {
		return errors.New("id cannot be empty")
	}

	if len(title) < 10 {
		return errors.New("title must be at least 3 characters long")
	}

	if len(description) < 10 {
		return errors.New("description must be at least 10 characters long")
	}

	if len(content) < 10 {
		return errors.New("content must be at least 10 characters long")
	}

	if !postValueObject.IsValidPostStatus(post.Status) {
		return errors.New("status is not valid")
	}

	// and more validations can be added here

	return nil
}
