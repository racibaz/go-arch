package ports

import (
	"context"

	"github.com/racibaz/go-arch/internal/modules/user/domain"
)

// UserRepository Secondary port: UserRepository interface
type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	Register(ctx context.Context, user *domain.User) error
	Me(ctx context.Context, id string) (*domain.User, error)
	IsExists(ctx context.Context, email string) (bool, error)
	UpdateWebUserRefreshToken(ctx context.Context, id, refreshToken string) error
	UpdateMobileUserRefreshToken(ctx context.Context, id, refreshToken string) error
}
