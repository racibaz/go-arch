package commands

import (
	"context"
	"fmt"

	applicationPorts "github.com/racibaz/go-arch/internal/modules/shared/application/ports"
	"github.com/racibaz/go-arch/internal/modules/user/domain/ports"
	"github.com/racibaz/go-arch/pkg/helper"
	"github.com/racibaz/go-arch/pkg/logger"
	"github.com/racibaz/go-arch/pkg/messaging"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type LogoutHandler struct {
	userRepository   ports.UserRepository
	logger           logger.Logger
	messagePublisher messaging.UserMessagePublisher
	tracer           trace.Tracer
}

// Ensure RegisterUserHandler implements the CommandHandler interface
var _ applicationPorts.CommandHandler[LogoutCommandV1] = (*LogoutHandler)(nil)

func NewLogoutHandler(
	userRepository ports.UserRepository,
	logger logger.Logger,
	messagePublisher messaging.UserMessagePublisher,
) *LogoutHandler {
	return &LogoutHandler{
		userRepository:   userRepository,
		logger:           logger,
		messagePublisher: messagePublisher,
		tracer:           otel.Tracer("LogoutHandler"),
	}
}

func (h LogoutHandler) Handle(ctx context.Context, cmd LogoutCommandV1) error {
	ctx, span := h.tracer.Start(ctx, "RegisterUser - Handler")
	defer span.End()

	user, getUserErr := h.userRepository.GetByID(ctx, cmd.UserID)
	if getUserErr != nil {
		h.logger.Error("User Not Found By ID: %v", getUserErr.Error())
		return getUserErr
	}

	// Save the new user to the repository
	savingErr := h.DeleteRefreshToken(ctx, cmd.UserID, cmd.Platform)

	if savingErr != nil {
		h.logger.Error("Error saving user: %v", savingErr)
		return savingErr
	}

	// Publish user logged out event
	if messageErr := h.messagePublisher.PublishUserLoggedOut(ctx, user); messageErr != nil {
		return fmt.Errorf("error publishing user logged out event: %w", messageErr)
	}

	h.logger.Info("User created successfully with ID: %s", user.ID())

	return nil
}

func (h LogoutHandler) DeleteRefreshToken(
	ctx context.Context,
	id, platform string,
) error {
	switch platform {
	case helper.PlatformWeb:
		return h.userRepository.DeleteWebUserRefreshToken(ctx, id)
	case helper.PlatformMobile:
		return h.userRepository.DeleteMobileUserRefreshToken(ctx, id)
	default:
		return fmt.Errorf("unknown platform: %s", platform)
	}
}
