// Package domain contains the domain logic for managing blog posts.

package domain

const (
	PostCreatedEvent = "posts.PostCreated"
	PostDeletedEvent = "post.PostDeleted"
)

type PostCreated struct {
	Post *Post
}

func (PostCreated) EventName() string { return PostCreatedEvent }

type PostDeleted struct {
	Post *Post
}

func (PostDeleted) EventName() string { return PostDeletedEvent }
