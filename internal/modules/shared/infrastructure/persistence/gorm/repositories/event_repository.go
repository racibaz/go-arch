package repositories

import (
	"sync"

	"github.com/racibaz/go-arch/internal/modules/shared/domain"
	"github.com/racibaz/go-arch/internal/modules/shared/domain/ports"
	"github.com/racibaz/go-arch/internal/modules/shared/infrastructure/persistence/gorm/entities"
	eventMapper "github.com/racibaz/go-arch/internal/modules/shared/infrastructure/persistence/gorm/mappers"
	"github.com/racibaz/go-arch/pkg/database"
	"gorm.io/gorm"
)

// GormEventRepository Secondary adapter: postgreSQL implementation
type GormEventRepository struct {
	DB *gorm.DB
	sync.Mutex
}

var _ ports.EventRepository = (*GormEventRepository)(nil)

func New() *GormEventRepository {
	return &GormEventRepository{
		DB: database.Connection(),
	}
}

func (repo *GormEventRepository) Save(event *domain.Event) error {
	var newEvent entities.Event

	persistenceModel := eventMapper.ToPersistence(*event)

	err := repo.DB.Create(&persistenceModel).Scan(&newEvent).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *GormEventRepository) GetByID(streamID string) (*domain.Event, error) {
	var event domain.Event

	if err := repo.DB.Where("stream_id = ?", streamID).First(&event).Error; err != nil {
		return nil, err
	}

	return &event, nil
}

func (repo *GormEventRepository) Update(event *domain.Event) error {
	// TODO implement me
	panic("implement me")
}

func (repo *GormEventRepository) Delete(id string) error {
	// TODO implement me
	panic("implement me")
}

func (repo *GormEventRepository) List() ([]*domain.Event, error) {
	// TODO implement me
	panic("implement me")
}

func (repo *GormEventRepository) IsExists(title, description string) (bool, error) {
	// TODO implement me
	panic("implement me")
}
