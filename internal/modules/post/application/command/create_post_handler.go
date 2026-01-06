package command

import (
	"context"
	"fmt"
	"time"

	applicationPorts "github.com/racibaz/go-arch/internal/modules/post/application/ports"
	"github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/internal/modules/post/domain/ports"
	"github.com/racibaz/go-arch/pkg/logger"
	"github.com/racibaz/go-arch/pkg/messaging"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// CreatePostHandler handles the creation of new posts.
type CreatePostHandler struct {
	PostRepository   ports.PostRepository
	logger           logger.Logger
	messagePublisher messaging.MessagePublisher
	tracer           trace.Tracer
}

// Ensure RemovePostHandler implements the CommandHandler interface
var _ applicationPorts.CommandHandler[CreatePostCommand] = (*CreatePostHandler)(nil)

// NewCreatePostHandler initializes a new CreatePostHandler with the provided PostRepository.
func NewCreatePostHandler(
	postRepository ports.PostRepository,
	logger logger.Logger,
	messagePublisher messaging.MessagePublisher,
) *CreatePostHandler {
	return &CreatePostHandler{
		PostRepository:   postRepository,
		logger:           logger,
		messagePublisher: messagePublisher,
		tracer:           otel.Tracer("CreatePostHandler"),
	}
}

// Handle processes the CreatePostCommand to create a new post.
func (h CreatePostHandler) Handle(ctx context.Context, cmd CreatePostCommand) error {
	ctx, span := h.tracer.Start(ctx, "CreatePost - Handler")
	defer span.End()

	// Create a new post using the factory
	post, _ := domain.Create(
		cmd.ID,
		cmd.UserID,
		cmd.Title,
		cmd.Description,
		cmd.Content,
		cmd.Status,
		time.Now(),
		time.Now(),
	)

	// check is the post exists in db?
	isExists, err := h.PostRepository.IsExists(ctx, post.Title, post.Description)
	if err != nil {
		h.logger.Error("Error saving post: %v", err.Error())
		return fmt.Errorf("error checking if post exists: %v", err)
	}

	// If the post already exists, return an error
	if isExists {
		h.logger.Info(
			"Post already exists with title: %s and description: %s",
			post.Title,
			post.Description,
		)
		return domain.ErrAlreadyExists
	}

	// Save the new post to the repository
	savingErr := h.PostRepository.Save(ctx, post)

	if savingErr != nil {
		h.logger.Error("Error saving post: %v", savingErr)
		return savingErr
	}

	// Publish an event indicating that a new post has been created
	if messageErr := h.messagePublisher.PublishPostCreated(ctx, post); messageErr != nil {
		return fmt.Errorf("failed to publish the post created event: %v", messageErr)
	}

	h.logger.Info("Post created successfully with ID: %s", post.ID())

	return nil
}

func (h CreatePostHandler) GetById(ctx context.Context, id string) (*domain.Post, error) {
	ctx, span := h.tracer.Start(ctx, "GetById - Handler")
	defer span.End()

	return h.PostRepository.GetByID(ctx, id)
}
