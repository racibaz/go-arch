package factories

import (
	"github.com/racibaz/go-arch/internal/modules/post/domain"
	postValueObject "github.com/racibaz/go-arch/internal/modules/post/domain/value_objects"
	"github.com/racibaz/go-arch/pkg/uuid"
)

// PostFactory is a factory for creating Post instances.
type PostFactory struct {
}

// New initializes a new PostFactory and returns a Post with default values.
func (pf *PostFactory) New() *domain.Post {

	// This function initializes a new PostFactory.
	// This factory method creates a new Post with default values.
	return &domain.Post{
		ID:          uuid.NewUuid().ToString(),
		Title:       "Default Title",
		Description: "Default Description",
		Status:      postValueObject.PostStatusDraft,
	}
}
