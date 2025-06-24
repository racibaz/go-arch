package domain

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	TitleMinLength       = 10
	DescriptionMinLength = 10
	ContentMinLength     = 10
)

var (
	ErrNotFound             = errors.New("the post was not found")
	ErrAlreadyExists        = errors.New("the post already exists")
	ErrEmptyId              = errors.New("id cannot be empty")
	ErrMinTitleLength       = errors.New(fmt.Sprintf("title must be at least %d characters long", TitleMinLength))
	ErrMinDescriptionLength = errors.New(fmt.Sprintf("description must be at least %d characters long", DescriptionMinLength))
	ErrMinContentLength     = errors.New(fmt.Sprintf("content must be at least %d characters long", ContentMinLength))
	ErrInvalidStatus        = errors.New("status is not valid")
)

type Post struct {
	ID          string
	Title       string
	Description string
	Content     string
	Status      PostStatus
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
		return ErrEmptyId
	}

	if len(post.Title) < TitleMinLength {
		return ErrMinTitleLength
	}

	if len(post.Description) < DescriptionMinLength {
		return ErrMinDescriptionLength
	}

	if len(post.Content) < ContentMinLength {
		return ErrMinContentLength
	}

	if !IsValidPostStatus(post.Status) {
		return ErrInvalidStatus
	}

	// and more validations can be added here

	return nil
}
