package commands

import (
	"context"

	applicationPorts "github.com/racibaz/go-arch/internal/modules/shared/application/ports"
	"github.com/racibaz/go-arch/internal/modules/user/domain"
	"github.com/racibaz/go-arch/internal/modules/user/domain/ports"
	"github.com/racibaz/go-arch/pkg/logger"
	"github.com/racibaz/go-arch/pkg/messaging"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// LoginHandler handles the login query.
type LoginHandler struct {
	UserRepository   ports.UserRepository
	logger           logger.Logger
	messagePublisher messaging.PostMessagePublisher
	tracer           trace.Tracer
	passwordHasher   ports.PasswordHasher
}

// Ensure LoginHandler implements the CommandHandler interface
var _ applicationPorts.CommandHandler[LoginCommandV1] = (*LoginHandler)(nil)

// NewLoginHandler creates a new instance of CreatePostHandler with the provided dependencies.
func NewLoginHandler(
	userRepository ports.UserRepository,
	logger logger.Logger,
	messagePublisher messaging.PostMessagePublisher,
	passwordHasher ports.PasswordHasher,
) *LoginHandler {
	return &LoginHandler{
		UserRepository:   userRepository,
		logger:           logger,
		messagePublisher: messagePublisher,
		tracer:           otel.Tracer("LoginHandler"),
		passwordHasher:   passwordHasher,
	}
}

func (h LoginHandler) Handle(ctx context.Context, cmd LoginCommandV1) error {
	ctx, span := h.tracer.Start(ctx, "Login - Handler")
	defer span.End()

	existingUser, getUserByEmailErr := h.UserRepository.GetUserByEmail(ctx, cmd.Email)
	if getUserByEmailErr != nil {
		h.logger.Error("User Not Found By Email: %v", getUserByEmailErr.Error())
		return getUserByEmailErr
	}

	hashedPassword, hashErr := h.passwordHasher.HashPassword(cmd.Password)
	if hashErr != nil {
		h.logger.Error("Error hashing password: %v", hashErr.Error())
		return hashErr
	}

	isCorrect := h.passwordHasher.VerifyPassword(cmd.Password, hashedPassword)

	if false == isCorrect {
		h.logger.Error("Invalid password for user with email: %s", cmd.Email)
		return domain.ErrInvalidCredentials
	}

	// todo check the status of the user (active, banned, etc.)

	/*	accessToken, generateJwtErr := helper.GenerateJWT(existingUser.ID(), existingUser.Name, platform)
		if generateJwtErr != nil {
			return generateJwtErr
		}*/

	// Log successful login
	h.logger.Info("User logged in successfully. Id: %s", existingUser.ID())

	return nil
}
