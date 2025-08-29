package commands

import (
	applicationPorts "github.com/racibaz/go-arch/internal/modules/post/application/ports"
	"github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/internal/modules/post/domain/ports"
	"github.com/racibaz/go-arch/pkg/logger"
	"time"
)

type CreatePostInput struct {
	ID          string // Unique identifier for the post
	UserID      string
	Title       string
	Description string
	Content     string
	Status      domain.PostStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time // ISO 8601 format
}

type CreatePostService struct {
	PostRepository ports.PostRepository
	logger         logger.Logger
}

var _ applicationPorts.PostService = (*CreatePostService)(nil)

// NewCreatePostService initializes a new CreatePostService with the provided PostRepository.
func NewCreatePostService(postRepository ports.PostRepository, logger logger.Logger) *CreatePostService {
	return &CreatePostService{
		PostRepository: postRepository,
		logger:         logger,
	}
}

func (postService CreatePostService) CreatePost(postInput CreatePostInput) error {

	// Create a new post using the factory
	post, _ := domain.Create(
		postInput.ID,
		postInput.UserID,
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

func (postService CreatePostService) GetById(id string) (*domain.Post, error) {

	return postService.PostRepository.GetByID(id)
}
