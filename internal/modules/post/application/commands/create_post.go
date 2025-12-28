package commands

import (
	"context"
	"fmt"
	"github.com/racibaz/go-arch/internal/modules/post/application/dtos"
	applicationPorts "github.com/racibaz/go-arch/internal/modules/post/application/ports"
	"github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/internal/modules/post/domain/ports"
	"github.com/racibaz/go-arch/pkg/logger"
	"github.com/racibaz/go-arch/pkg/messaging"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"time"
)

type CreatePostService struct {
	PostRepository   ports.PostRepository
	logger           logger.Logger
	messagePublisher messaging.MessagePublisher
	tracer           trace.Tracer
}

var _ applicationPorts.PostService = (*CreatePostService)(nil)

// NewCreatePostService initializes a new CreatePostService with the provided PostRepository.
func NewCreatePostService(postRepository ports.PostRepository, logger logger.Logger, messagePublisher messaging.MessagePublisher) *CreatePostService {
	return &CreatePostService{
		PostRepository:   postRepository,
		logger:           logger,
		messagePublisher: messagePublisher,
		tracer:           otel.Tracer("CreatePostService"),
	}
}

func (postService CreatePostService) CreatePost(ctx context.Context, postInput dto.CreatePostInput) error {

	ctx, span := postService.tracer.Start(ctx, "CreatePost - Service")
	defer span.End()

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

	//todo the err msg have to come from this function
	// check is the post exists in db?
	isExists, err := postService.PostRepository.IsExists(ctx, post.Title, post.Description)

	//todo when the check, we can check bool and err together
	if err != nil {
		return err
	}
	//todo when the check, we can check bool and err together
	// If the post already exists, return an error
	if isExists {
		postService.logger.Info("Post already exists with title: %s and description: %s", post.Title, post.Description)
		return domain.ErrAlreadyExists
	}

	savingErr := postService.PostRepository.Save(ctx, post)

	if savingErr != nil {
		postService.logger.Error("Error saving post: %v", savingErr)
		return savingErr
	}
	// Publish an event indicating that a new post has been created
	if messageErr := postService.messagePublisher.PublishPostCreated(ctx, post); messageErr != nil {
		return fmt.Errorf("failed to publish the post created event: %v", messageErr)
	}

	postService.logger.Info("Post created successfully with ID: %s", post.ID())

	return nil
}

func (postService CreatePostService) GetById(ctx context.Context, id string) (*domain.Post, error) {

	ctx, span := postService.tracer.Start(ctx, "GetById - Service")
	defer span.End()

	return postService.PostRepository.GetByID(ctx, id)
}
