package in_memory

import (
	"errors"
	"fmt"
	"github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/pkg/uuid"
	"sync"
)

// InMemoryPostRepository Secondary adapter: InMemory implementation
type InMemoryPostRepository struct {
	posts map[uuid.Uuid]*domain.Post
	sync.Mutex
}

func New() *InMemoryPostRepository {
	return &InMemoryPostRepository{
		posts: make(map[uuid.Uuid]*domain.Post),
	}
}

// TODO it should return mapper or aggregate root
func (pr *InMemoryPostRepository) Save(post *domain.Post) error {
	pr.Mutex.Lock()
	defer pr.Mutex.Unlock()

	fmt.Printf("Creating post with ID %s\n", post)
	// Validate the post ID
	if post.ID == "" {
		return errors.New("post ID cannot be empty")
	}

	parsedUUID, err := uuid.Parse(post.ID)

	if _, exists := pr.posts[parsedUUID]; exists {
		return errors.New("post ID cannot be empty")
	}

	if err != nil {
		return errors.New("invalid Post ID format")
	}

	fmt.Printf("Post with ID %s created successfully\n", post.ID)

	return nil
}

func (pr *InMemoryPostRepository) GetByID(id string) (*domain.Post, error) {

	postID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid Post ID format")
	}

	if id == "" {
		return nil, errors.New("post ID cannot be nil")
	}
	//TODO post gelmiyor slice olabilir..

	exists := pr.posts[postID]

	if exists == nil {
		return nil, errors.New("post not found")
	}

	return exists, nil
}

func (pr *InMemoryPostRepository) Update(post *domain.Post) error {
	//TODO implement me
	panic("implement me")
}

func (pr *InMemoryPostRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func (pr *InMemoryPostRepository) List() ([]*domain.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *InMemoryPostRepository) IsExists(title, description string) (bool, error) {

	//TODO implement me
	return false, nil
}
