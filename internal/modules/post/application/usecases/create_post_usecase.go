package usecases

import (
	useCaseInputs "github.com/racibaz/go-arch/internal/modules/post/application/usecases/inputs"
	"github.com/racibaz/go-arch/internal/modules/post/domain"
	postFactory "github.com/racibaz/go-arch/internal/modules/post/domain/factories"
	"github.com/racibaz/go-arch/internal/modules/post/domain/ports"
	"github.com/racibaz/go-arch/pkg/logger"
	"time"
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

	// Create a new post using the factory
	post, _ := postFactory.New(
		postInput.ID,
		postInput.Title,
		postInput.Description,
		postInput.Content,
		postInput.Status,
		time.Now(),
		time.Now(),
	)

	// check is the post exists in db?
	isExists, err := postService.PostRepository.IsExists(post.Title, post.Description)

	if err != nil {
		return err
	}

	// If the post already exists, return an error
	if isExists {
		logger.Info("Post already exists")
		return domain.ErrAlreadyExists
	}

	savingErr := postService.PostRepository.Save(post)

	if savingErr != nil {
		logger.Error("Error saving post")
		return savingErr
	}

	logger.Info("Post created successfully")

	return nil
}

func (postService CreatePostUseCase) GetById(id string) (*domain.Post, error) {

	return postService.PostRepository.GetByID(id)
}
