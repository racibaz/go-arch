package queries

import (
	"context"
	"github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/internal/modules/post/domain/ports"
	"github.com/racibaz/go-arch/pkg/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type GetPostInput struct {
	ID     string // Unique identifier for the post
	UserID string
}

type GetPostService struct {
	PostRepository ports.PostRepository
	logger         logger.Logger
	tracer         trace.Tracer
}

// NewGetPostService initializes a new GetPostService with the provided PostRepository.
func NewGetPostService(postRepository ports.PostRepository, logger logger.Logger) *GetPostService {
	return &GetPostService{
		PostRepository: postRepository,
		logger:         logger,
		tracer:         otel.Tracer("GetPostService"),
	}
}

func (postService GetPostService) GetPostByID(ctx context.Context, postInput GetPostInput) (*domain.Post, error) {

	ctx, span := postService.tracer.Start(ctx, "GetById - Service")
	defer span.End()

	post, err := postService.PostRepository.GetByID(ctx, postInput.ID)
	if err != nil {
		return nil, err
	}

	return post, err
}
