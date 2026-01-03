package commands

import (
	"github.com/racibaz/go-arch/internal/modules/post/domain/ports"
	"github.com/racibaz/go-arch/pkg/logger"
)

type RemovePostInput struct {
	ID     string // Unique identifier for the post
	UserID string
}

type RemovePostService struct {
	PostRepository ports.PostRepository
	logger         logger.Logger
}

// var _ applicationPorts.PostService = (*RemovePostService)(nil)

// NewRemovePostService initializes a new CreatePostService with the provided PostRepository.
func NewRemovePostService(
	postRepository ports.PostRepository,
	logger logger.Logger,
) *RemovePostService {
	return &RemovePostService{
		PostRepository: postRepository,
		logger:         logger,
	}
}

func (postService CreatePostService) Remove(postInput RemovePostInput) error {
	// todo implement me

	return nil
}
