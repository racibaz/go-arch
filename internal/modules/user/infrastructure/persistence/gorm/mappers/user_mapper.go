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
		Aggregate: es.NewAggregate(userEntity.ID, domain.UserAggregate),
		UserName:  userEntity.UserName,
		Email:     userEntity.Email,
		Password:  userEntity.Password,
		Status:    status,
		CreatedAt: userEntity.CreatedAt,
		UpdatedAt: userEntity.UpdatedAt,
	}, nil
}

// ToPersistence maps a domain User to a persistence entity User
func ToPersistence(user *domain.User) (*entity.User, error) {
	if user == nil {
		return nil, errors.New("user is nil")
	}

	return &entity.User{
		ID:        user.ID(),
		UserName:  user.UserName,
		Email:     user.Email,
		Password:  user.Password,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
