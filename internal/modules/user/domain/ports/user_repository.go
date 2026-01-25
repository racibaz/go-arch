package ports

import (
	"context"

	"github.com/racibaz/go-arch/internal/modules/user/domain"
)

// UserRepository Secondary port: UserRepository interface
type UserRepository interface {
	Login(ctx context.Context, data any) error
	Register(ctx context.Context, data any) (*domain.User, error)
	Me(ctx context.Context, id string) (*domain.User, error)
}
