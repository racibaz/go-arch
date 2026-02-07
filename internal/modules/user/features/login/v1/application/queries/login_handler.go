package query

import (
	"context"
	"fmt"

	applicationPorts "github.com/racibaz/go-arch/internal/modules/shared/application/ports"
	"github.com/racibaz/go-arch/internal/modules/user/domain"
	"github.com/racibaz/go-arch/internal/modules/user/domain/ports"
	"github.com/racibaz/go-arch/pkg/helper"
	"github.com/racibaz/go-arch/pkg/logger"
	"github.com/racibaz/go-arch/pkg/messaging"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// LoginHandler handles the login query.
type LoginHandler struct {
	userRepository   ports.UserRepository
	logger           logger.Logger
	messagePublisher messaging.UserMessagePublisher
	tracer           trace.Tracer
	passwordHasher   ports.PasswordHasher
}

// Ensure LoginHandler implements the CommandHandler interface
var _ applicationPorts.QueryHandler[LoginQueryV1, *LoginQueryResponse] = (*LoginHandler)(nil)

// NewLoginHandler creates a new instance of CreatePostHandler with the provided dependencies.
func NewLoginHandler(
	userRepository ports.UserRepository,
	logger logger.Logger,
	messagePublisher messaging.UserMessagePublisher,
	passwordHasher ports.PasswordHasher,
) *LoginHandler {
	return &LoginHandler{
		userRepository:   userRepository,
		logger:           logger,
		messagePublisher: messagePublisher,
		tracer:           otel.Tracer("LoginHandler"),
		passwordHasher:   passwordHasher,
	}
}

func (h LoginHandler) Handle(ctx context.Context, cmd LoginQueryV1) (*LoginQueryResponse, error) {
	ctx, span := h.tracer.Start(ctx, "Login - Handler")
	defer span.End()

	existingUser, getUserByEmailErr := h.userRepository.GetByEmail(ctx, cmd.Email)
	if getUserByEmailErr != nil {
		h.logger.Error("User Not Found By Email: %v", getUserByEmailErr.Error())
		return nil, getUserByEmailErr
	}

	// todo check user status

	hashedPassword, hashErr := h.passwordHasher.HashPassword(cmd.Password)
	if hashErr != nil {
		h.logger.Error("Error hashing password: %v", hashErr.Error())
		return nil, hashErr
	}

	isCorrect := h.passwordHasher.VerifyPassword(cmd.Password, hashedPassword)

	if false == isCorrect {
		h.logger.Error("Invalid password for user with email: %s", cmd.Email)
		return nil, domain.ErrInvalidCredentials
	}

	// todo check the status of the user (active, banned, etc.)

	accessToken, generateJwtErr := helper.GenerateJWT(
		existingUser.ID(),
		existingUser.Name,
		cmd.Platform,
	)
	if generateJwtErr != nil {
		return nil, generateJwtErr
	}

	refreshToken, refreshTokenErr := helper.GenerateRefreshToken()
	if refreshTokenErr != nil {
		return nil, fmt.Errorf(
			"login failed - could not generate refresh token: %w",
			refreshTokenErr,
		)
	}

	updateRefreshTokenErr := h.UpdateRefreshToken(
		ctx,
		existingUser.ID(),
		refreshToken,
		cmd.Platform,
	)
	if updateRefreshTokenErr != nil {
		return nil, fmt.Errorf(
			"login failed - could not update user refresh token: %w",
			updateRefreshTokenErr,
		)
	}

	// Publish user logged in event
	if messageErr := h.messagePublisher.PublishUserLoggedIn(ctx, existingUser); messageErr != nil {
		return nil, fmt.Errorf("failed to publish the user logged in event: %v", messageErr)
	}

	// Log successful login
	h.logger.Info("User logged in successfully. Id: %s", existingUser.ID())

	return &LoginQueryResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       existingUser.ID(),
	}, nil
}

func (h LoginHandler) UpdateRefreshToken(
	ctx context.Context,
	id, refreshToken, platform string,
) error {
	switch platform {
	case helper.PlatformWeb:
		return h.userRepository.UpdateWebUserRefreshToken(ctx, id, refreshToken)
	case helper.PlatformMobile:
		return h.userRepository.UpdateMobileUserRefreshToken(ctx, id, refreshToken)
	default:
		return fmt.Errorf("invalid platform: %s", platform)
	}
}
