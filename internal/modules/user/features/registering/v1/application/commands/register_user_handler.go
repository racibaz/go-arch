package commands

import (
	"context"
	"fmt"
	"time"

	applicationPorts "github.com/racibaz/go-arch/internal/modules/shared/application/ports"
	"github.com/racibaz/go-arch/internal/modules/user/domain"
	"github.com/racibaz/go-arch/internal/modules/user/domain/ports"
	"github.com/racibaz/go-arch/pkg/logger"
	"github.com/racibaz/go-arch/pkg/messaging"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type RegisterUserHandler struct {
	UserRepository   ports.UserRepository
	logger           logger.Logger
	messagePublisher messaging.UserMessagePublisher
	tracer           trace.Tracer
	passwordHasher   ports.PasswordHasher
}

// Ensure RegisterUserHandler implements the CommandHandler interface
var _ applicationPorts.CommandHandler[RegisterUserCommandV1] = (*RegisterUserHandler)(nil)

func NewRegisterUserHandler(
	userRepository ports.UserRepository,
	logger logger.Logger,
	messagePublisher messaging.UserMessagePublisher,
	passwordHasher ports.PasswordHasher,
) *RegisterUserHandler {
	return &RegisterUserHandler{
		UserRepository:   userRepository,
		logger:           logger,
		messagePublisher: messagePublisher,
		tracer:           otel.Tracer("RegisterUserHandler"),
		passwordHasher:   passwordHasher,
	}
}

func (h RegisterUserHandler) Handle(ctx context.Context, cmd RegisterUserCommandV1) error {
	ctx, span := h.tracer.Start(ctx, "RegisterUser - Handler")
	defer span.End()

	hashedPassword, hashErr := h.passwordHasher.HashPassword(cmd.Password)
	if hashErr != nil {
		h.logger.Error("Error hashing password: %v", hashErr.Error())
		return fmt.Errorf("error hashing password: %v", hashErr)
	}
	// Create a new user using the factory
	user, err := domain.Create(
		cmd.ID,
		cmd.UserName,
		cmd.Email,
		hashedPassword,
		domain.StatusDraft,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		h.logger.Error("Error creating user: %v", err.Error())
		return fmt.Errorf("error creating user: %v", err)
	}

	// check is the user exists in db?
	isExists, isExistsErr := h.UserRepository.IsExists(ctx, user.Email)
	if isExistsErr != nil {
		h.logger.Error("Error saving user: %v", isExistsErr.Error())
		return fmt.Errorf("error checking if user exists: %v", isExistsErr)
	}

	// If the user already exists, return an error
	if isExists {
		h.logger.Info(
			"User already exists with email: %s",
			user.Email,
		)
		return domain.ErrAlreadyExists
	}

	// Save the new user to the repository
	savingErr := h.UserRepository.Register(ctx, user)

	if savingErr != nil {
		h.logger.Error("Error saving user: %v", savingErr)
		return savingErr
	}

	// Publish an event indicating that a new user has been created
	if messageErr := h.messagePublisher.PublishUserRegistered(ctx, user); messageErr != nil {
		return fmt.Errorf("failed to publish the user created event: %v", messageErr)
	}

	h.logger.Info("User created successfully with ID: %s", user.ID())

	return nil
}
