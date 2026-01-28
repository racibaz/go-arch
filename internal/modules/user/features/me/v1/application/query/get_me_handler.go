package query

import (
	"context"

	"github.com/racibaz/go-arch/internal/modules/user/domain/ports"
	"github.com/racibaz/go-arch/pkg/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// GetMeHandler handles the retrieval of the current user's information.
type GetMeHandler struct {
	UserRepository ports.UserRepository
	logger         logger.Logger
	tracer         trace.Tracer
}

// NewGetMeHandler creates a new instance of GetMeHandler.
func NewGetMeHandler(userRepository ports.UserRepository, logger logger.Logger) *GetMeHandler {
	return &GetMeHandler{
		UserRepository: userRepository,
		logger:         logger,
		tracer:         otel.Tracer("GetMeHandler"),
	}
}

// Handle processes the GetMeByIdQuery and returns the user's information.
func (h *GetMeHandler) Handle(
	ctx context.Context,
	query GetMeByIdQuery,
) (GetMeByIdQueryResponse, error) {
	ctx, span := h.tracer.Start(ctx, "GetMe - Handler")
	defer span.End()

	user, err := h.UserRepository.Me(ctx, query.ID)
	if err != nil {
		return GetMeByIdQueryResponse{}, err
	}

	// Map domain.User to GetMeByIdQueryResponse
	userView := GetMeByIdQueryResponse{
		ID:        user.ID(),
		Name:      user.Name,
		Email:     user.Email,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return userView, err
}
