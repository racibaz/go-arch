package queries

import (
	"github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/internal/modules/post/domain/ports"
	"github.com/racibaz/go-arch/pkg/logger"
)

type GetPostInput struct {
	ID     string // Unique identifier for the post
	UserID string
}

type GetPostService struct {
	PostRepository ports.PostRepository
	logger         logger.Logger
}

//var _ applicationPorts.PostService = (*RemovePostService)(nil)

// NewGetPostService initializes a new GetPostService with the provided PostRepository.
func NewGetPostService(postRepository ports.PostRepository, logger logger.Logger) *GetPostService {
	return &GetPostService{
		PostRepository: postRepository,
		logger:         logger,
	}
}

// todo add userID check
// todo reponse type should be DTO
func (postService GetPostService) GetPostByID(postInput GetPostInput) (*domain.Post, error) {

	post, err := postService.PostRepository.GetByID(postInput.ID)
	if err != nil {
		return nil, err
	}

	return post, err
}
