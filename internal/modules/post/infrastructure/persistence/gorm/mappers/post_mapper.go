package mappers

import (
	domain "github.com/racibaz/go-arch/internal/modules/post/domain"
	entity "github.com/racibaz/go-arch/internal/modules/post/infrastructure/persistence/gorm/entities"
	"github.com/racibaz/go-arch/pkg/es"
)

func ToDomain(postEntity entity.Post) domain.Post {

	status := domain.PostStatus(postEntity.Status)

	return domain.Post{
		Aggregate:   es.NewAggregate(postEntity.ID, domain.PostAggregate),
		UserID:      postEntity.UserID,
		Title:       postEntity.Title,
		Description: postEntity.Description,
		Content:     postEntity.Content,
		Status:      status,
		CreatedAt:   postEntity.CreatedAt,
		UpdatedAt:   postEntity.UpdatedAt,
	}

}

func ToPersistence(post domain.Post) entity.Post {

	return entity.Post{
		ID:          post.ID(),
		UserID:      post.UserID,
		Title:       post.Title,
		Description: post.Description,
		Content:     post.Content,
		Status:      int(post.Status),
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
	}

}
