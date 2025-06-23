package domain

import (
	"errors"
	postValueObject "github.com/racibaz/go-arch/internal/modules/post/domain/value_objects"
	"strings"
	"time"
)

var (
	TitleMinLength       = 10
	DescriptionMinLength = 10
	ContentMinLength     = 10
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

func (post *Post) Sanitize() {

	// Trim whitespace from the input parameters
	post.ID = strings.TrimSpace(post.ID)
	post.Title = strings.TrimSpace(post.Title)
	post.Description = strings.TrimSpace(post.Description)
	post.Content = strings.TrimSpace(post.Content)
}

// Validate checks if the Post fields are valid.
func (post *Post) Validate() error {

	// Sanitize the input parameters
	post.Sanitize()

	// Validate the input parameters
	if post.ID == "" {
		return errors.New("id cannot be empty")
	}

	if len(post.Title) < TitleMinLength {
		return errors.New("title must be at least 3 characters long")
	}

	if len(post.Description) < DescriptionMinLength {
		return errors.New("description must be at least 10 characters long")
	}

	if len(post.Content) < ContentMinLength {
		return errors.New("content must be at least 10 characters long")
	}

	if !postValueObject.IsValidPostStatus(post.Status) {
		return errors.New("status is not valid")
	}

	// and more validations can be added here

	return nil
}
