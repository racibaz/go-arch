package mappers

import (
	domain "github.com/racibaz/go-arch/internal/modules/post/domain"
	valueObject "github.com/racibaz/go-arch/internal/modules/post/domain/value_objects"
	entity "github.com/racibaz/go-arch/internal/modules/post/infrastructure/persistence/gorm/entities"
)

func ToDomain(postEntity entity.Post) domain.Post {

	status := valueObject.PostStatus(postEntity.Status)

	return domain.Post{
		ID:          postEntity.ID,
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
		ID:          post.ID,
		Title:       post.Title,
		Description: post.Description,
		Content:     post.Content,
		Status:      int(post.Status),
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
	}

}
