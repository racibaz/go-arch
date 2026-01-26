package ports

import (
	"context"

	"github.com/racibaz/go-arch/internal/modules/user/domain"
)

type LoginData struct {
	Email    string
	Password string
}

// UserRepository Secondary port: UserRepository interface
type UserRepository interface {
	Login(ctx context.Context, data LoginData) (*domain.User, error)
	Register(ctx context.Context, user *domain.User) error
	Me(ctx context.Context, id string) (*domain.User, error)
	IsExists(ctx context.Context, email string) (bool, error)
}
