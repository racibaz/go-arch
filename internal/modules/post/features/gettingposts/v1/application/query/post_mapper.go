package query

import (
	"errors"

	domain "github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/pkg/es"
)

func ToDomains(entities []*PostView) ([]*domain.Post, error) {
	if entities == nil {
		return nil, errors.New("post entity is nil")
	}

	posts := make([]*domain.Post, 0)

	for _, entity := range entities {

		status := domain.PostStatus(entity.Status)

		post := &domain.Post{
			Aggregate:   es.NewAggregate(entity.ID, domain.PostAggregate),
			UserID:      entity.UserID,
			Title:       entity.Title,
			Description: entity.Description,
			Content:     entity.Content,
			Status:      status,
			CreatedAt:   entity.CreatedAt,
			UpdatedAt:   entity.UpdatedAt,
		}

		posts = append(posts, post)

	}

	return posts, nil
}

// ToDto maps domain.Post slices to PostView slices
func ToDto(posts []*domain.Post) ([]*PostView, error) {
	if posts == nil {
		return nil, errors.New("posts are nil")
	}

	postViews := make([]*PostView, 0)

	for _, post := range posts {

		postView := &PostView{
			ID:          post.ID(),
			UserID:      post.UserID,
			Title:       post.Title,
			Description: post.Description,
			Content:     post.Content,
			Status:      int(post.Status),
			CreatedAt:   post.CreatedAt,
			UpdatedAt:   post.UpdatedAt,
		}
		postViews = append(postViews, postView)

	}
	return postViews, nil
}
