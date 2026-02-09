package mappers

import (
	"errors"

	domain "github.com/racibaz/go-arch/internal/modules/user/domain"
	entity "github.com/racibaz/go-arch/internal/modules/user/infrastructure/persistence/gorm/entities"
	"github.com/racibaz/go-arch/pkg/es"
)

func ToDomain(userEntity *entity.User) (*domain.User, error) {
	if userEntity == nil {
		return nil, errors.New("user entity is nil")
	}

	status := domain.UserStatus(userEntity.Status)

	return &domain.User{
		Aggregate:            es.NewAggregate(userEntity.ID, domain.UserAggregate),
		Name:                 userEntity.Name,
		Email:                userEntity.Email,
		Password:             userEntity.Password,
		Status:               status,
		RefreshTokenWeb:      userEntity.RefreshTokenWeb,
		RefreshTokenWebAt:    userEntity.RefreshTokenWebAt,
		RefreshTokenMobile:   userEntity.RefreshTokenMobile,
		RefreshTokenMobileAt: userEntity.RefreshTokenMobileAt,
		CreatedAt:            userEntity.CreatedAt,
		UpdatedAt:            userEntity.UpdatedAt,
	}, nil
}

// ToPersistence maps a domain User to a persistence entity User
func ToPersistence(user *domain.User) (*entity.User, error) {
	if user == nil {
		return nil, errors.New("user is nil")
	}

	return &entity.User{
		ID:                   user.ID(),
		Name:                 user.Name,
		Email:                user.Email,
		Password:             user.Password,
		Status:               user.Status,
		RefreshTokenWeb:      user.RefreshTokenWeb,
		RefreshTokenWebAt:    user.RefreshTokenWebAt,
		RefreshTokenMobile:   user.RefreshTokenMobile,
		RefreshTokenMobileAt: user.RefreshTokenMobileAt,
		CreatedAt:            user.CreatedAt,
		UpdatedAt:            user.UpdatedAt,
	}, nil
}
