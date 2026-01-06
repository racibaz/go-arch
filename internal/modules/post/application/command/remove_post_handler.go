package command

import (
	"context"

	applicationPorts "github.com/racibaz/go-arch/internal/modules/post/application/ports"
	"github.com/racibaz/go-arch/internal/modules/post/domain/ports"
	"github.com/racibaz/go-arch/pkg/logger"
)

// RemovePostHandler handles the removal of posts.
type RemovePostHandler struct {
	PostRepository ports.PostRepository
	logger         logger.Logger
}

// Ensure RemovePostHandler implements the CommandHandler interface
var _ applicationPorts.CommandHandler[RemovePostCommand] = (*RemovePostHandler)(nil)

// NewRemovePostHandler initializes a new CreatePostHandler with the provided PostRepository.
func NewRemovePostHandler(
	postRepository ports.PostRepository,
	logger logger.Logger,
) *RemovePostHandler {
	return &RemovePostHandler{
		PostRepository: postRepository,
		logger:         logger,
	}
}

// Handle processes the RemovePostCommand to remove a post.
func (h *RemovePostHandler) Handle(ctx context.Context, cmd RemovePostCommand) error {
	// todo implement me
	return nil
}
