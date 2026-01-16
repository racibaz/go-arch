package in_memory

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/internal/modules/post/domain/ports"
	"github.com/racibaz/go-arch/pkg/uuid"
)

// Repository Secondary adapters: InMemory implementation
type Repository struct {
	posts map[uuid.Uuid]*domain.Post
	sync.Mutex
}

var _ ports.PostRepository = (*Repository)(nil)

func New() *Repository {
	return &Repository{
		posts: make(map[uuid.Uuid]*domain.Post),
	}
}

// TODO it should return mapper or aggregate root
func (pr *Repository) Save(ctx context.Context, post *domain.Post) error {
	pr.Mutex.Lock()
	defer pr.Mutex.Unlock()

	fmt.Printf("Creating post with ID %s\n", post.ID())
	// Validate the post ID
	if post.ID() == "" {
		return errors.New("post ID cannot be empty")
	}

	parsedUUID, err := uuid.Parse(post.ID())

	if _, exists := pr.posts[parsedUUID]; exists {
		return errors.New("post ID cannot be empty")
	}

	if err != nil {
		return errors.New("invalid Post ID format")
	}

	fmt.Printf("Post with ID %s created successfully\n", post.ID())

	return nil
}

func (pr *Repository) GetByID(ctx context.Context, id string) (*domain.Post, error) {
	postID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid Post ID format")
	}

	if id == "" {
		return nil, errors.New("post ID cannot be nil")
	}

	exists := pr.posts[postID]

	if exists == nil {
		return nil, errors.New("post not found")
	}

	return exists, nil
}

func (pr *Repository) Update(ctx context.Context, post *domain.Post) error {
	// TODO implement me
	panic("implement me")
}

func (pr *Repository) Delete(ctx context.Context, id string) error {
	// TODO implement me
	panic("implement me")
}

func (pr *Repository) List(ctx context.Context) ([]*domain.Post, error) {
	// TODO implement me
	panic("implement me")
}

func (repo *Repository) IsExists(ctx context.Context, title, description string) (bool, error) {
	// TODO implement me
	return false, nil
}
