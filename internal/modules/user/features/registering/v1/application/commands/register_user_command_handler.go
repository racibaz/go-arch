package commands

import (
	"context"
	applicationPorts "github.com/racibaz/go-arch/internal/modules/shared/application/ports"
	"github.com/racibaz/go-arch/internal/modules/user/domain/ports"
	"github.com/racibaz/go-arch/pkg/logger"
	"github.com/racibaz/go-arch/pkg/messaging"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type RegisterUserHandler struct {
	UserRepository   ports.UserRepository
	logger           logger.Logger
	messagePublisher messaging.MessagePublisher
	tracer           trace.Tracer
}

// Ensure LoginHandler implements the CommandHandler interface
var _ applicationPorts.CommandHandler[RegisterUserCommandV1] = (*RegisterUserHandler)(nil)

func NewRegisterUserHandler(
	userRepository ports.UserRepository,
	logger logger.Logger,
	messagePublisher messaging.MessagePublisher,
) *RegisterUserHandler {
	return &RegisterUserHandler{
		UserRepository:   userRepository,
		logger:           logger,
		messagePublisher: messagePublisher,
		tracer:           otel.Tracer("LoginHandler"),
	}
}

func (h RegisterUserHandler) Handle(ctx context.Context, cmd RegisterUserCommandV1) error {
	ctx, span := h.tracer.Start(ctx, "RegisterUser - Handler")
	defer span.End()

	return nil
}
