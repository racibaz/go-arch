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

// LoginHandler handles the login commands.
type LoginHandler struct {
	UserRepository   ports.UserRepository
	logger           logger.Logger
	messagePublisher messaging.PostMessagePublisher
	tracer           trace.Tracer
}

// Ensure LoginHandler implements the CommandHandler interface
var _ applicationPorts.CommandHandler[LoginCommandV1] = (*LoginHandler)(nil)

// NewLoginHandler creates a new instance of CreatePostHandler with the provided dependencies.
func NewLoginHandler(
	userRepository ports.UserRepository,
	logger logger.Logger,
	messagePublisher messaging.PostMessagePublisher,
) *LoginHandler {
	return &LoginHandler{
		UserRepository:   userRepository,
		logger:           logger,
		messagePublisher: messagePublisher,
		tracer:           otel.Tracer("LoginHandler"),
	}
}

func (h LoginHandler) Handle(ctx context.Context, cmd LoginCommandV1) error {
	ctx, span := h.tracer.Start(ctx, "Login - Handler")
	defer span.End()

	return nil
}
