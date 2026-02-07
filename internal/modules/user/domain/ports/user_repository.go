package ports

import (
	"context"

	"github.com/racibaz/go-arch/internal/modules/user/domain"
)

// UserRepository Secondary port: UserRepository interface
type UserRepository interface {
	GetByID(ctx context.Context, id string) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	Register(ctx context.Context, user *domain.User) error
	Me(ctx context.Context, refreshToken string) (*domain.User, error)
	IsExists(ctx context.Context, email string) (bool, error)
	UpdateWebUserRefreshToken(ctx context.Context, id, refreshToken string) error
	UpdateMobileUserRefreshToken(ctx context.Context, id, refreshToken string) error
	DeleteWebUserRefreshToken(ctx context.Context, id string) error
	DeleteMobileUserRefreshToken(ctx context.Context, id string) error
	GetUserByRefreshTokenAtWeb(ctx context.Context, refreshToken string) (*domain.User, error)
	GetUserByRefreshTokenAtMobile(ctx context.Context, refreshToken string) (*domain.User, error)
}
