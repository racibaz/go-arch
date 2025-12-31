package repositories

import (
	"context"
	"testing"
	"time"

	"github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/internal/modules/post/infrastructure/persistence/gorm/entities"
	postMapper "github.com/racibaz/go-arch/internal/modules/post/infrastructure/persistence/gorm/mappers"
	"github.com/racibaz/go-arch/pkg/uuid"
	"github.com/stretchr/testify/assert"
)

func createTestPost() (*domain.Post, error) {
	id := uuid.NewUuid().ToString()
	userID := uuid.NewUuid().ToString()
	now := time.Now()

	return domain.Create(
		id,
		userID,
		"Test Title That Is Long Enough",
		"Test Description That Is Long Enough",
		"Test Content That Is Long Enough",
		domain.PostStatusDraft,
		now,
		now,
	)
}

func TestNew(t *testing.T) {
	// Test constructor - this will use the actual database connection
	// In a real test environment, you'd want to mock this
	// For now, just test that it doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("NewGormPostRepository() panicked: %v", r)
		}
	}()

	repo := NewGormPostRepository()
	assert.NotNil(t, repo)
	// Note: DB might be nil if database is not connected in test environment
}

func TestGormPostRepository_Update_Panics(t *testing.T) {
	repo := NewGormPostRepository()

	post, err := createTestPost()
	assert.NoError(t, err)

	ctx := context.Background()

	assert.Panics(t, func() {
		repo.Update(ctx, post)
	})
}

func TestGormPostRepository_Delete_Panics(t *testing.T) {
	repo := NewGormPostRepository()

	ctx := context.Background()

	assert.Panics(t, func() {
		repo.Delete(ctx, uuid.NewUuid().ToString())
	})
}

func TestGormPostRepository_List_Panics(t *testing.T) {
	repo := NewGormPostRepository()

	ctx := context.Background()

	assert.Panics(t, func() {
		repo.List(ctx)
	})
}

func TestPostMapperRoundTrip(t *testing.T) {
	// Test that the mapper round-trip works correctly
	post, err := createTestPost()
	assert.NoError(t, err)

	// Convert to persistence model
	persistenceModel, err := postMapper.ToPersistence(post)
	assert.NoError(t, err)

	// Convert back to domain model
	domainModel, err := postMapper.ToDomain(persistenceModel)
	assert.NoError(t, err)

	// Verify the round-trip conversion
	assert.Equal(t, post.ID(), domainModel.ID())
	assert.Equal(t, post.UserID, domainModel.UserID)
	assert.Equal(t, post.Title, domainModel.Title)
	assert.Equal(t, post.Description, domainModel.Description)
	assert.Equal(t, post.Content, domainModel.Content)
	assert.Equal(t, post.Status, domainModel.Status)
	assert.Equal(t, post.CreatedAt, domainModel.CreatedAt)
	assert.Equal(t, post.UpdatedAt, domainModel.UpdatedAt)
}

func TestPostMapperToPersistence(t *testing.T) {
	post, err := createTestPost()
	assert.NoError(t, err)

	entity, err := postMapper.ToPersistence(post)
	assert.NoError(t, err)

	assert.Equal(t, post.ID(), entity.ID)
	assert.Equal(t, post.UserID, entity.UserID)
	assert.Equal(t, post.Title, entity.Title)
	assert.Equal(t, post.Description, entity.Description)
	assert.Equal(t, post.Content, entity.Content)
	assert.Equal(t, int(post.Status), entity.Status)
}

func TestPostMapperToDomain(t *testing.T) {
	entity := entities.Post{
		ID:          uuid.NewUuid().ToString(),
		UserID:      uuid.NewUuid().ToString(),
		Title:       "Test Title",
		Description: "Test Description",
		Content:     "Test Content",
		Status:      int(domain.PostStatusPublished),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	domainPost, err := postMapper.ToDomain(&entity)
	assert.NoError(t, err)

	assert.Equal(t, entity.ID, domainPost.ID())
	assert.Equal(t, entity.UserID, domainPost.UserID)
	assert.Equal(t, entity.Title, domainPost.Title)
	assert.Equal(t, entity.Description, domainPost.Description)
	assert.Equal(t, entity.Content, domainPost.Content)
	assert.Equal(t, domain.PostStatus(entity.Status), domainPost.Status)
}

func TestRepositoryImplementsInterface(t *testing.T) {
	// Test that GormPostRepository implements the PostRepository interface
	var repo interface{} = &GormPostRepository{}
	_, ok := repo.(interface {
		Save(ctx context.Context, post *domain.Post) error
		GetByID(ctx context.Context, id string) (*domain.Post, error)
		Update(ctx context.Context, post *domain.Post) error
		Delete(ctx context.Context, id string) error
		List(ctx context.Context) ([]*domain.Post, error)
		IsExists(ctx context.Context, title, description string) (bool, error)
	})

	assert.True(t, ok, "GormPostRepository should implement PostRepository interface")
}
