package in_memory

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/racibaz/go-arch/internal/modules/user/domain"
	"github.com/racibaz/go-arch/internal/modules/user/domain/ports"
	"github.com/racibaz/go-arch/pkg/uuid"
)

// InMemoryRepository Secondary adapters: InMemory implementation
type InMemoryRepository struct {
	users map[uuid.Uuid]*domain.User
	sync.Mutex
}

var _ ports.UserRepository = (*InMemoryRepository)(nil)

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		users: make(map[uuid.Uuid]*domain.User),
	}
}

// TODO it should return mapper or aggregate root
func (pr *InMemoryRepository) Save(ctx context.Context, user *domain.User) error {
	pr.Mutex.Lock()
	defer pr.Mutex.Unlock()

	fmt.Printf("Creating user with ID %s\n", user.ID())
	// Validate the post ID
	if user.ID() == "" {
		return errors.New("user ID cannot be empty")
	}

	parsedUUID, err := uuid.Parse(user.ID())

	if _, exists := pr.users[parsedUUID]; exists {
		return errors.New("user ID cannot be empty")
	}

	if err != nil {
		return errors.New("invalid User ID format")
	}

	fmt.Printf("User with ID %s created successfully\n", user.ID())

	return nil
}

func (pr *InMemoryRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	userID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid User ID format")
	}

	if id == "" {
		return nil, errors.New("user ID cannot be nil")
	}

	exists := pr.users[userID]

	if exists == nil {
		return nil, errors.New("user not found")
	}

	return exists, nil
}

func (pr *InMemoryRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	// TODO implement me
	panic("implement me")
}

func (pr *InMemoryRepository) Register(ctx context.Context, user *domain.User) error {
	// TODO implement me
	panic("implement me")
}

func (pr *InMemoryRepository) Me(ctx context.Context, refreshToken string) (*domain.User, error) {
	// TODO implement me
	panic("implement me")
}

func (pr *InMemoryRepository) IsExists(ctx context.Context, email string) (bool, error) {
	// TODO implement me
	panic("implement me")
}

func (pr *InMemoryRepository) UpdateWebUserRefreshToken(
	ctx context.Context,
	id, refreshToken string,
) error {
	// TODO implement me
	panic("implement me")
}

func (pr *InMemoryRepository) UpdateMobileUserRefreshToken(
	ctx context.Context,
	id, refreshToken string,
) error {
	// TODO implement me
	panic("implement me")
}

func (pr *InMemoryRepository) DeleteWebUserRefreshToken(ctx context.Context, id string) error {
	// TODO implement me
	panic("implement me")
}

func (pr *InMemoryRepository) DeleteMobileUserRefreshToken(ctx context.Context, id string) error {
	// TODO implement me
	panic("implement me")
}

func (pr *InMemoryRepository) GetUserByRefreshTokenAtWeb(
	ctx context.Context,
	refreshToken string,
) (*domain.User, error) {
	// TODO implement me
	panic("implement me")
}

func (pr *InMemoryRepository) GetUserByRefreshTokenAtMobile(
	ctx context.Context,
	refreshToken string,
) (*domain.User, error) {
	// TODO implement me
	panic("implement me")
}
