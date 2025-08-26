package usecases

import (
	applicationPorts "github.com/racibaz/go-arch/internal/modules/post/application/ports"
	useCaseInputs "github.com/racibaz/go-arch/internal/modules/post/application/usecases/inputs"
	"github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/internal/modules/post/domain/ports"
	"github.com/racibaz/go-arch/pkg/logger"
	"time"
)

type PostUseCase struct {
	PostRepository ports.PostRepository
	logger         logger.Logger
}

var _ applicationPorts.PostService = (*PostUseCase)(nil)

// NewPostUseCase initializes a new PostUseCase with the provided PostRepository.
func NewPostUseCase(postRepository ports.PostRepository, logger logger.Logger) *PostUseCase {
	return &PostUseCase{
		PostRepository: postRepository,
		logger:         logger,
	}
}

func (postService PostUseCase) CreatePost(postInput useCaseInputs.CreatePostInput) error {

	// Create a new post using the factory
	post, _ := domain.Create(
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
		postService.logger.Info("Post already exists with title: %s and description: %s", post.Title, post.Description)
		return domain.ErrAlreadyExists
	}

	savingErr := postService.PostRepository.Save(post)

	if savingErr != nil {
		postService.logger.Error("Error saving post: %v", savingErr)
		return savingErr
	}

	postService.logger.Info("Post created successfully with ID: %s", post.ID)

	return nil
}

func (postService PostUseCase) GetById(id string) (*domain.Post, error) {

	return postService.PostRepository.GetByID(id)
}
