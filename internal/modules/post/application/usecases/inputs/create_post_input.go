package inputs

import (
	postValueObject "github.com/racibaz/go-arch/internal/modules/post/domain/value_objects"
	"time"
)

type CreatePostInput struct {
	ID          string // Unique identifier for the post
	Title       string
	Description string
	Content     string
	Status      postValueObject.PostStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time // ISO 8601 format
}
