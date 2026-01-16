package transformers

import (
	"github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/internal/modules/post/features/gettingpostbyid/v1/adapters/endpoints/dtos"
	"github.com/racibaz/go-arch/internal/modules/post/features/gettingpostbyid/v1/application/query"
)

func FromPostCoreToHTTP(input *domain.Post) *dtos.Post {
	if input == nil {
		return nil
	}

	result := &dtos.Post{
		Title:       input.Title,
		Description: input.Description,
		Content:     input.Content,
		Status:      input.Status.String(),
	}

	return result
}

func FromPostViewToHTTP(input *query.GetPostByIdQueryResponse) *dtos.Post {
	if input == nil {
		return nil
	}

	result := &dtos.Post{
		Title:       input.Title,
		Description: input.Description,
		Content:     input.Content,
		Status:      domain.PostStatus(input.Status).String(),
	}

	return result
}
