package mappers

import (
	"errors"

	domain "github.com/racibaz/go-arch/internal/modules/post/domain"
	entity "github.com/racibaz/go-arch/internal/modules/post/infrastructure/persistence/gorm/entities"
	"github.com/racibaz/go-arch/pkg/es"
)

// ToDomain maps a persistence entity Post to a domain Post
func ToDomain(postEntity *entity.Post) (*domain.Post, error) {
	if postEntity == nil {
		return nil, errors.New("post entity is nil")
	}

	status := domain.PostStatus(postEntity.Status)

	return &domain.Post{
		Aggregate:   es.NewAggregate(postEntity.ID, domain.PostAggregate),
		UserID:      postEntity.UserID,
		Title:       postEntity.Title,
		Description: postEntity.Description,
		Content:     postEntity.Content,
		Status:      status,
		CreatedAt:   postEntity.CreatedAt,
		UpdatedAt:   postEntity.UpdatedAt,
	}, nil
}

// ToPersistence maps a domain Post to a persistence entity Post
func ToPersistence(post *domain.Post) (*entity.Post, error) {
	if post == nil {
		return nil, errors.New("post is nil")
	}

	return &entity.Post{
		ID:          post.ID(),
		UserID:      post.UserID,
		Title:       post.Title,
		Description: post.Description,
		Content:     post.Content,
		Status:      post.Status,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
	}, nil
}
