package dto

import (
	"github.com/racibaz/go-arch/internal/modules/post/domain"
	"time"
)

type CreatePostInput struct {
	ID          string // Unique identifier for the post
	UserID      string
	Title       string
	Description string
	Content     string
	Status      domain.PostStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time // ISO 8601 format
}
