package domain

import (
	"errors"
	"fmt"
	"github.com/racibaz/go-arch/pkg/ddd"
	"github.com/racibaz/go-arch/pkg/es"
	"strings"
	"time"
)

const PostAggregate = "posts.Post"

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

var _ interface {
	es.EventApplier
	es.Snapshotter
} = (*Post)(nil)

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

func Create(id, userID, title, description, content string, status PostStatus, createdAt, updatedAt time.Time) (*Post, error) {

	// This factory method creates a new Post with default values if you want.
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

	//validate the post before returning it
	err := post.Validate()
	if err != nil {
		return nil, err
	}

	post.AddEvent(PostCreatedEvent, &PostCreated{
		Post: post,
	})

	return post, nil
}

func (p *Post) Delete() {
	//todo implement me
}

func (p *Post) Sanitize() {

	// Trim whitespace from the input parameters
	p.Title = strings.TrimSpace(p.Title)
	p.Description = strings.TrimSpace(p.Description)
	p.Content = strings.TrimSpace(p.Content)
}

// Validate checks if the Post fields are valid.
func (p *Post) Validate() error {

	// Sanitize the input parameters
	p.Sanitize()

	// Validate the input parameters
	if p.Aggregate.ID() == "" {
		return ErrEmptyId
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

// ApplyEvent implements es.EventApplier
func (p *Post) ApplyEvent(event ddd.Event) error {
	switch payload := event.Payload().(type) {
	case *PostCreated:
		p.UserID = payload.Post.UserID
		p.Status = payload.Post.Status
	default:
		return errors.New(fmt.Sprintf("%T received the event %s with unexpected payload %T", p, event.EventName(), payload))
	}

	return nil
}

// ApplySnapshot implements es.Snapshotter
func (p *Post) ApplySnapshot(snapshot es.Snapshot) error {
	switch ss := snapshot.(type) {
	case *PostV1:
		p.UserID = ss.UserID
		p.Title = ss.Title
		p.Content = ss.Content
		p.Status = ss.Status

	default:
		return errors.New(fmt.Sprintf("%T received the unexpected snapshot %T", p, snapshot))
	}

	return nil
}

// ToSnapshot implements es.Snapshotter
func (s Post) ToSnapshot() es.Snapshot {
	return PostV1{
		UserID:  s.UserID,
		Title:   s.Title,
		Content: s.Content,
		Status:  s.Status,
	}
}
