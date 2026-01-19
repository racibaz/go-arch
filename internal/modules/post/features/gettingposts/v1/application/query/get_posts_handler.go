package query

import (
	"context"

	"github.com/racibaz/go-arch/internal/modules/post/domain/ports"
	"github.com/racibaz/go-arch/pkg/helper"
	"github.com/racibaz/go-arch/pkg/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type GetPostsHandler struct {
	PostRepository ports.PostRepository
	logger         logger.Logger
	tracer         trace.Tracer
}

// NewGetPostsHandler initializes a new GetPostHandler with the provided PostRepository.
func NewGetPostsHandler(
	postRepository ports.PostRepository,
	logger logger.Logger,
) *GetPostsHandler {
	return &GetPostsHandler{
		PostRepository: postRepository,
		logger:         logger,
		tracer:         otel.Tracer("GetPostsHandler"),
	}
}

// Handle retrieves a post by its ID using the PostRepository.
func (h *GetPostsHandler) Handle(
	ctx context.Context,
	query helper.Pagination,
) (GetPostsQueryResponse, error) {
	ctx, span := h.tracer.Start(ctx, "GetPosts - Handler")
	defer span.End()

	posts, err := h.PostRepository.List(ctx, query)
	if err != nil {
		return GetPostsQueryResponse{}, err
	}

	postView, _ := ToDto(posts)
	responseDtO := GetPostsQueryResponse{postView}

	return responseDtO, err
}
