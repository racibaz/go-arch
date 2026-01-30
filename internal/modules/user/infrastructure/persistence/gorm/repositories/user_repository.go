package repositories

import (
	"context"
	"fmt"
	"sync"

	"github.com/racibaz/go-arch/internal/modules/user/domain"
	"github.com/racibaz/go-arch/internal/modules/user/domain/ports"
	"github.com/racibaz/go-arch/internal/modules/user/infrastructure/persistence/gorm/entities"
	userMapper "github.com/racibaz/go-arch/internal/modules/user/infrastructure/persistence/gorm/mappers"
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

// NewGormUserRepository creates a new instance of GormUserRepository
func NewGormUserRepository() *GormUserRepository {
	return &GormUserRepository{
		DB: database.Connection(),
	}
}

// Save persists a new user in the database
func (repo *GormUserRepository) Save(ctx context.Context, user *domain.User) error {
	var newUser entities.User

	persistenceModel, persistenceErr := userMapper.ToPersistence(user)
	if persistenceErr != nil {
		return fmt.Errorf("failed to map user to persistence model: %w", persistenceErr)
	}

	err := repo.DB.WithContext(ctx).Create(&persistenceModel).Scan(&newUser).Error
	if err != nil {
		return fmt.Errorf("new user creation is failed: %w", err)
	}

	return nil
}

// GetByID retrieves a user by its ID from the database
func (repo *GormUserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User

	if err := repo.DB.WithContext(ctx).
		Where("id = ?", id).
		First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Update modifies an existing user in the database
func (repo *GormUserRepository) Update(ctx context.Context, user *domain.User) error {
	err := repo.DB.WithContext(ctx).Updates(user)
	if err != nil {
		return err.Error
	}
	return nil
}

// Delete removes a user from the database by its ID
func (repo *GormUserRepository) Delete(ctx context.Context, id string) error {
	err := repo.DB.WithContext(ctx).Delete(&domain.User{}, "id = ?", id)
	if err != nil {
		return err.Error
	}

	return nil
}

// List retrieves a list of users with pagination support
func (repo *GormUserRepository) List(
	ctx context.Context,
	pagination helper.Pagination,
) ([]*domain.User, error) {
	var users []*domain.User

	err := repo.DB.WithContext(ctx).
		Scopes(helper.Paginate(pagination)).
		Find(&users).
		Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

// IsExists checks if a user exists by email
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

	if user.Email != "" {
		return true, nil
	}

	return false, nil
}

// Register registers a new user in the database
func (repo *GormUserRepository) Register(ctx context.Context, user *domain.User) error {
	var newUser entities.User

	persistenceModel, persistenceErr := userMapper.ToPersistence(user)
	if persistenceErr != nil {
		return fmt.Errorf("failed to map user to persistence model: %w", persistenceErr)
	}

	err := repo.DB.WithContext(ctx).Create(&persistenceModel).Scan(&newUser).Error
	if err != nil {
		return fmt.Errorf("new user creation is failed: %w", err)
	}

	return nil
}

// Me retrieves the current user's information by ID
func (repo *GormUserRepository) Me(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User

	if err := repo.DB.WithContext(ctx).
		Where("id = ?", id).
		First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *GormUserRepository) GetByEmail(
	ctx context.Context,
	email string,
) (*domain.User, error) {
	var userEntity entities.User

	if err := repo.DB.WithContext(ctx).
		Where("email = ?", email).
		First(&userEntity).Error; err != nil {
		return nil, err
	}

	user, mapErr := userMapper.ToDomain(&userEntity)
	if mapErr != nil {
		return nil, fmt.Errorf("failed to map user entity to domain model: %w", mapErr)
	}

	return user, nil
}

func (repo *GormUserRepository) UpdateWebUserRefreshToken(
	ctx context.Context,
	id string,
	refreshToken string,
) error {
	err := repo.DB.WithContext(ctx).
		Model(&domain.User{}).
		Where("id::text = ?", id).
		Update("refresh_token_web", refreshToken).
		Update("refresh_token_web_at", gorm.Expr("NOW()")).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *GormUserRepository) UpdateMobileUserRefreshToken(
	ctx context.Context,
	id string,
	refreshToken string,
) error {
	err := repo.DB.WithContext(ctx).
		Model(&domain.User{}).
		Where("id::text = ?", id).
		Update("refresh_token_mobile", refreshToken).
		Update("refresh_token_mobile_at", gorm.Expr("NOW()")).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *GormUserRepository) DeleteWebUserRefreshToken(
	ctx context.Context,
	id string,
) error {
	err := repo.DB.WithContext(ctx).
		Model(&domain.User{}).
		Where("id::text = ?", id).
		Update("refresh_token_web", gorm.Expr("null")).
		Update("refresh_token_web_at", gorm.Expr("null")).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *GormUserRepository) DeleteMobileUserRefreshToken(
	ctx context.Context,
	id string,
) error {
	err := repo.DB.WithContext(ctx).
		Model(&domain.User{}).
		Where("id::text = ?", id).
		Update("refresh_token_mobile", gorm.Expr("null")).
		Update("refresh_token_mobile_at", gorm.Expr("null")).Error
	if err != nil {
		return err
	}
	return nil
}
