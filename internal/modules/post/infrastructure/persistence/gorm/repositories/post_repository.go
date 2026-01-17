package repositories

import (
	"context"
	"fmt"
	"sync"

	"github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/internal/modules/post/domain/ports"
	"github.com/racibaz/go-arch/internal/modules/post/infrastructure/persistence/gorm/entities"
	postMapper "github.com/racibaz/go-arch/internal/modules/post/infrastructure/persistence/gorm/mappers"
	"github.com/racibaz/go-arch/pkg/database"
	"github.com/racibaz/go-arch/pkg/helper"
	"gorm.io/gorm"
)

// GormPostRepository Secondary adapters: PostgreSQL implementation
type GormPostRepository struct {
	DB *gorm.DB
	sync.Mutex
}

var _ ports.PostRepository = (*GormPostRepository)(nil)

func NewGormPostRepository() *GormPostRepository {
	return &GormPostRepository{
		DB: database.Connection(),
	}
}

func (repo *GormPostRepository) Save(ctx context.Context, post *domain.Post) error {
	var newPost entities.Post

	persistenceModel, persistenceErr := postMapper.ToPersistence(post)
	if persistenceErr != nil {
		return fmt.Errorf("failed to map post to persistence model: %w", persistenceErr)
	}

	err := repo.DB.WithContext(ctx).Create(&persistenceModel).Scan(&newPost).Error
	if err != nil {
		return fmt.Errorf("new post creation is failed: %w", err)
	}

	return nil
}

func (repo *GormPostRepository) GetByID(ctx context.Context, id string) (*domain.Post, error) {
	var post domain.Post

	if err := repo.DB.WithContext(ctx).
		Where("id = ?", id).
		First(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (repo *GormPostRepository) Update(ctx context.Context, post *domain.Post) error {
	err := repo.DB.WithContext(ctx).Updates(post)
	if err != nil {
		return err.Error
	}
	return nil
}

func (repo *GormPostRepository) Delete(ctx context.Context, id string) error {
	err := repo.DB.WithContext(ctx).Delete(&domain.Post{}, "id = ?", id)
	if err != nil {
		return err.Error
	}

	return nil
}

func (repo *GormPostRepository) List(
	ctx context.Context,
	pagination helper.Pagination,
) ([]*domain.Post, error) {
	var posts []*domain.Post

	err := repo.DB.WithContext(ctx).Scopes(helper.Paginate(pagination)).Find(&posts)
	if err != nil {
		return nil, err.Error
	}

	return posts, nil
}

func (repo *GormPostRepository) IsExists(
	ctx context.Context,
	title, description string,
) (bool, error) {
	var post domain.Post

	if err := repo.DB.WithContext(ctx).
		Where("title = ?", title).
		Where("description = ?", description).
		First(&post).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	if post.ID() != "" {
		return true, nil
	}

	return false, nil
}
