package usecases

import (
	useCaseInputs "github.com/racibaz/go-arch/internal/modules/post/application/usecases/inputs"
	"github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/internal/modules/post/domain/ports"
)

type CreatePostUseCase struct {
	PostRepository ports.PostRepository
}

// NewCreatetPostUseCase initializes a new CreatePostUseCase with the provided PostRepository.
func NewCreatetPostUseCase(postRepository ports.PostRepository) *CreatePostUseCase {
	return &CreatePostUseCase{
		PostRepository: postRepository,
	}
}

func (postService CreatePostUseCase) CreatePost(postInput useCaseInputs.CreatePostInput) error {
	err := postService.PostRepository.Create(&domain.Post{
		ID:          postInput.ID,
		Title:       postInput.Title,
		Description: postInput.Description,
		Status:      postInput.Status,
		CreatedAt:   postInput.CreatedAt,
		UpdatedAt:   postInput.UpdatedAt,
	})

	if err != nil {
		return err
	}

	return nil
}

func (postService CreatePostUseCase) GetById(id string) (*domain.Post, error) {

	return postService.PostRepository.GetByID(id)
}
