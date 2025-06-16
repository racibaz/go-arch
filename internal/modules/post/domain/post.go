package domain

import (
	postValueObject "github.com/racibaz/go-arch/internal/modules/post/domain/value-objects"
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
