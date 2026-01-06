package query

import (
	"context"

	"github.com/racibaz/go-arch/internal/modules/post/domain/ports"
	"github.com/racibaz/go-arch/pkg/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type GetPostHandler struct {
	PostRepository ports.PostRepository
	logger         logger.Logger
	tracer         trace.Tracer
}

// NewGetPostHandler initializes a new GetPostHandler with the provided PostRepository.
func NewGetPostHandler(postRepository ports.PostRepository, logger logger.Logger) *GetPostHandler {
	return &GetPostHandler{
		PostRepository: postRepository,
		logger:         logger,
		tracer:         otel.Tracer("GetPostHandler"),
	}
}

// Handle retrieves a post by its ID using the PostRepository.
func (h *GetPostHandler) Handle(
	ctx context.Context,
	query GetPostQuery,
) (PostView, error) {
	ctx, span := h.tracer.Start(ctx, "GetById - Handler")
	defer span.End()

	post, err := h.PostRepository.GetByID(ctx, query.ID)
	if err != nil {
		return PostView{}, err
	}

	// Map domain.Post to PostView DTO
	postView := PostView{
		ID:          post.ID(),
		UserID:      post.UserID,
		Title:       post.Title,
		Description: post.Description,
		Content:     post.Content,
		Status:      int(post.Status),
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
	}

	return postView, err
}
