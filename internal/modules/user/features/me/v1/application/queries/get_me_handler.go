package queries

import (
	"context"

	"github.com/racibaz/go-arch/internal/modules/user/domain/ports"
	"github.com/racibaz/go-arch/pkg/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// MeQueryHandler handles the retrieval of the current user's information.
type MeQueryHandler struct {
	userRepository ports.UserRepository
	logger         logger.Logger
	tracer         trace.Tracer
}

// NewMeQueryHandler creates a new instance of MeQueryHandler.
func NewMeQueryHandler(userRepository ports.UserRepository, logger logger.Logger) *MeQueryHandler {
	return &MeQueryHandler{
		userRepository: userRepository,
		logger:         logger,
		tracer:         otel.Tracer("MeQueryHandler"),
	}
}

// Handle processes the GetMeByIdQuery and returns the user's information.
func (h *MeQueryHandler) Handle(
	ctx context.Context,
	query MeQueryHandlerQuery,
) (*MeQueryHandlerResponse, error) {
	ctx, span := h.tracer.Start(ctx, "Me - Handler")
	defer span.End()

	user, err := h.userRepository.Me(ctx, query.RefreshToken)
	if err != nil {
		return nil, err
	}

	// Map the user domain model to the response DTO
	responsePayload := MeQueryHandlerResponse{
		ID:        user.ID(),
		Name:      user.Name,
		Email:     user.Email,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return &responsePayload, err
}
