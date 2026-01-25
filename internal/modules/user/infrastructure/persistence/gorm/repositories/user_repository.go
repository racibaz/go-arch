package repositories

import (
	"context"
	"fmt"
	userMapper "github.com/racibaz/go-arch/internal/modules/user/infrastructure/persistence/gorm/mappers"
	"sync"

	"github.com/racibaz/go-arch/internal/modules/user/domain"
	"github.com/racibaz/go-arch/internal/modules/user/domain/ports"
	"github.com/racibaz/go-arch/internal/modules/user/infrastructure/persistence/gorm/entities"
	"github.com/racibaz/go-arch/pkg/database"
	"github.com/racibaz/go-arch/pkg/helper"
	"gorm.io/gorm"
)

// GormUserRepository Secondary adapters: PostgreSQL implementation
type GormUserRepository struct {
	DB *gorm.DB
	sync.Mutex
}

var _ ports.UserRepository = (*GormUserRepository)(nil)

func NewGormUserRepository() *GormUserRepository {
	return &GormUserRepository{
		DB: database.Connection(),
	}
}

func (repo *GormUserRepository) Save(ctx context.Context, user *domain.User) error {
	var newUser entities.User

	persistenceModel, persistenceErr := userMapper.ToPersistence(user)
	if persistenceErr != nil {
		return fmt.Errorf("failed to map post to persistence model: %w", persistenceErr)
	}

	err := repo.DB.WithContext(ctx).Create(&persistenceModel).Scan(&newUser).Error
	if err != nil {
		return fmt.Errorf("new post creation is failed: %w", err)
	}

	return nil
}

func (repo *GormUserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	var post domain.User

	if err := repo.DB.WithContext(ctx).
		Where("id = ?", id).
		First(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (repo *GormUserRepository) Update(ctx context.Context, user *domain.User) error {
	err := repo.DB.WithContext(ctx).Updates(user)
	if err != nil {
		return err.Error
	}
	return nil
}

func (repo *GormUserRepository) Delete(ctx context.Context, id string) error {
	err := repo.DB.WithContext(ctx).Delete(&domain.User{}, "id = ?", id)
	if err != nil {
		return err.Error
	}

	return nil
}

func (repo *GormUserRepository) List(
	ctx context.Context,
	pagination helper.Pagination,
) ([]*domain.User, error) {
	var posts []*domain.User

	err := repo.DB.WithContext(ctx).
		Scopes(helper.Paginate(pagination)).
		Find(&posts).
		Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (repo *GormUserRepository) IsExists(
	ctx context.Context,
	email string,
) (bool, error) {
	var user domain.User

	if err := repo.DB.WithContext(ctx).
		Where("email = ?", email).
		First(&user).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	if user.ID() != "" {
		return true, nil
	}

	return false, nil
}

func (repo *GormUserRepository) Login(ctx context.Context, data any) error {
	//TODO implement me
	panic("implement me")
}

func (repo *GormUserRepository) Register(ctx context.Context, data any) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *GormUserRepository) Me(ctx context.Context, id string) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}
