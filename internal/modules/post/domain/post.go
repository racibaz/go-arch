// Package post contains the domain logic for managing blog posts.

package domain

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/racibaz/go-arch/pkg/es"
)

const (
	PostAggregate = "posts.Post"
)

var (
	TitleMinLength       = 10
	DescriptionMinLength = 10
	ContentMinLength     = 10
)

var (
	ErrNotFound       = errors.New("the post was not found")
	ErrAlreadyExists  = errors.New("the post already exists")
	ErrEmptyId        = errors.New("id cannot be empty")
	ErrEmptyUserId    = errors.New("user id cannot be empty")
	ErrMinTitleLength = errors.New(
		fmt.Sprintf("title must be at least %d characters long", TitleMinLength),
	)
	ErrMinDescriptionLength = errors.New(
		fmt.Sprintf("description must be at least %d characters long", DescriptionMinLength),
	)
	ErrMinContentLength = errors.New(
		fmt.Sprintf("content must be at least %d characters long", ContentMinLength),
	)
	ErrInvalidStatus = errors.New("status is not valid")
)

type Post struct {
	es.Aggregate
	UserID      string
	Title       string
	Description string
	Content     string
	Status      PostStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Create This factory method creates a new Post with default values if you want.
func Create(
	id, userID, title, description, content string,
	status PostStatus,
	createdAt, updatedAt time.Time,
) (*Post, error) {
	post := &Post{
		Aggregate:   es.NewAggregate(id, PostAggregate),
		UserID:      userID,
		Title:       title,
		Description: description,
		Content:     content,
		Status:      status,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	// validate the post before returning it
	err := post.validate()
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (p *Post) Delete() {
	// todo implement me
}

func (p *Post) sanitize() {
	// Trim whitespace from the input parameters
	p.UserID = strings.TrimSpace(p.UserID)
	p.Title = strings.TrimSpace(p.Title)
	p.Description = strings.TrimSpace(p.Description)
	p.Content = strings.TrimSpace(p.Content)
}

// Validate checks if the Post fields are valid.
func (p *Post) validate() error {
	// Sanitize the input parameters
	p.sanitize()

	// Validate the input parameters
	if p.Aggregate.ID() == "" {
		return ErrEmptyId
	}

	if p.UserID == "" {
		return ErrEmptyUserId
	}

	if len(p.Title) < TitleMinLength {
		return ErrMinTitleLength
	}

	if len(p.Description) < DescriptionMinLength {
		return ErrMinDescriptionLength
	}

	if len(p.Content) < ContentMinLength {
		return ErrMinContentLength
	}

	if !IsValidPostStatus(p.Status) {
		return ErrInvalidStatus
	}

	// and more validations can be added here

	return nil
}
