package presentation

import (
	"github.com/racibaz/go-arch/internal/modules/post/application/query"
	"github.com/racibaz/go-arch/internal/modules/post/domain"
)

func FromPostCoreToHTTP(input *domain.Post) *Post {
	if input == nil {
		return nil
	}

	result := &Post{
		Title:       input.Title,
		Description: input.Description,
		Content:     input.Content,
		Status:      input.Status.String(),
	}

	return result
}

func FromPostViewToHTTP(input *query.PostView) *Post {
	if input == nil {
		return nil
	}

	result := &Post{
		Title:       input.Title,
		Description: input.Description,
		Content:     input.Content,
		Status:      domain.PostStatus(input.Status).String(),
	}

	return result
}
