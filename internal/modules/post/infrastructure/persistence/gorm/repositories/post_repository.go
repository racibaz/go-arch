package repositories

import (
	"github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/internal/modules/post/domain/ports"
	"github.com/racibaz/go-arch/internal/modules/post/infrastructure/persistence/gorm/entities"
	postMapper "github.com/racibaz/go-arch/internal/modules/post/infrastructure/persistence/gorm/mappers"
	"github.com/racibaz/go-arch/pkg/database"
	"gorm.io/gorm"
	"sync"
)

// GormPostRepository Secondary adapter: PostgreSQL implementation
type GormPostRepository struct {
	DB *gorm.DB
	sync.Mutex
}

var _ ports.PostRepository = (*GormPostRepository)(nil)

func New() *GormPostRepository {
	return &GormPostRepository{
		DB: database.Connection(),
	}
}

func (repo *GormPostRepository) Save(post *domain.Post) error {
	var newPost entities.Post

	persistenceModel := postMapper.ToPersistence(*post)

	err := repo.DB.Create(&persistenceModel).Scan(&newPost).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *GormPostRepository) GetByID(id string) (*domain.Post, error) {

	var post domain.Post

	if err := repo.DB.Where("id = ?", id).First(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (repo *GormPostRepository) Update(post *domain.Post) error {
	//TODO implement me
	panic("implement me")
}

func (repo *GormPostRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func (repo *GormPostRepository) List() ([]*domain.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *GormPostRepository) IsExists(title, description string) (bool, error) {

	var post domain.Post

	repo.DB.Where("title = ?", title).Where("description = ?", description).First(&post)

	if post.ID() != "" {
		return true, nil
	}

	return false, nil
}
