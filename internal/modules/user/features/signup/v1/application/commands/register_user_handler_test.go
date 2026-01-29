package commands

import (
	"context"
	"errors"
	"fmt"
	"testing"

	sharedPorts "github.com/racibaz/go-arch/internal/modules/shared/application/ports"
	"github.com/racibaz/go-arch/internal/modules/user/domain"
	userDomainPortsMocks "github.com/racibaz/go-arch/internal/modules/user/testing/mocks/domain/ports"
	loggerMocks "github.com/racibaz/go-arch/pkg/logger"
	messagingMocks "github.com/racibaz/go-arch/pkg/messaging"
	"github.com/racibaz/go-arch/pkg/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewRegisterUserHandler(t *testing.T) {
	// Given
	mockRepo := userDomainPortsMocks.NewMockUserRepository(t)
	mockLogger := loggerMocks.NewMockLogger(t)
	mockPublisher := messagingMocks.NewMockUserMessagePublisher(t)
	mockHasher := userDomainPortsMocks.NewMockPasswordHasher(t)

	// When
	handler := NewRegisterUserHandler(mockRepo, mockLogger, mockPublisher, mockHasher)

	// Then
	assert.NotNil(t, handler)
	assert.Equal(t, mockRepo, handler.UserRepository)
	assert.Equal(t, mockLogger, handler.logger)
	assert.Equal(t, mockPublisher, handler.messagePublisher)
	assert.Equal(t, mockHasher, handler.passwordHasher)
	assert.NotNil(t, handler.tracer)
}

func TestRegisterUserHandler_ImplementsCommandHandlerInterface(t *testing.T) {
	// Given
	mockRepo := userDomainPortsMocks.NewMockUserRepository(t)
	mockLogger := loggerMocks.NewMockLogger(t)
	mockPublisher := messagingMocks.NewMockUserMessagePublisher(t)
	mockHasher := userDomainPortsMocks.NewMockPasswordHasher(t)
	handler := NewRegisterUserHandler(mockRepo, mockLogger, mockPublisher, mockHasher)

	// When & Then
	var _ sharedPorts.CommandHandler[RegisterUserCommandV1] = handler
}

func TestRegisterUserHandler_Handle(t *testing.T) {
	testCases := []struct {
		name        string
		command     RegisterUserCommandV1
		setupMocks  func(*userDomainPortsMocks.MockUserRepository, *loggerMocks.MockLogger, *messagingMocks.MockUserMessagePublisher, *userDomainPortsMocks.MockPasswordHasher)
		expectedErr error
	}{
		{
			name: "successful user registration",
			command: RegisterUserCommandV1{
				ID:       uuid.NewUuid().ToString(),
				Name:     "testuser",
				Email:    "test@example.com",
				Password: "password123",
			},
			setupMocks: func(mockRepo *userDomainPortsMocks.MockUserRepository, mockLogger *loggerMocks.MockLogger, mockPublisher *messagingMocks.MockUserMessagePublisher, mockHasher *userDomainPortsMocks.MockPasswordHasher) {
				mockHasher.EXPECT().HashPassword(mock.AnythingOfType("string")).Return("hashedpassword", nil).Once()
				mockRepo.EXPECT().IsExists(mock.Anything, mock.AnythingOfType("string")).Return(false, nil).Once()
				mockRepo.EXPECT().Register(mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()
				mockPublisher.EXPECT().PublishUserRegistered(mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()
				mockLogger.EXPECT().Info(mock.Anything, mock.Anything).Maybe()
				mockLogger.EXPECT().Error(mock.Anything, mock.Anything).Maybe()
			},
			expectedErr: nil,
		},
		{
			name: "password hashing error",
			command: RegisterUserCommandV1{
				ID:       uuid.NewUuid().ToString(),
				Name:     "testuser",
				Email:    "test@example.com",
				Password: "password123",
			},
			setupMocks: func(mockRepo *userDomainPortsMocks.MockUserRepository, mockLogger *loggerMocks.MockLogger, mockPublisher *messagingMocks.MockUserMessagePublisher, mockHasher *userDomainPortsMocks.MockPasswordHasher) {
				mockHasher.EXPECT().HashPassword(mock.AnythingOfType("string")).Return("", errors.New("hashing failed")).Once()
				mockLogger.EXPECT().Error(mock.Anything, mock.Anything).Maybe()
			},
			expectedErr: fmt.Errorf("error hashing password: %w", errors.New("hashing failed")),
		},
		{
			name: "domain creation error - invalid email",
			command: RegisterUserCommandV1{
				ID:       uuid.NewUuid().ToString(),
				Name:     "testuser",
				Email:    "invalid-email", // This will cause domain.Create to fail
				Password: "password123",
			},
			setupMocks: func(mockRepo *userDomainPortsMocks.MockUserRepository, mockLogger *loggerMocks.MockLogger, mockPublisher *messagingMocks.MockUserMessagePublisher, mockHasher *userDomainPortsMocks.MockPasswordHasher) {
				mockHasher.EXPECT().HashPassword(mock.AnythingOfType("string")).Return("hashedpassword", nil).Once()
				mockLogger.EXPECT().Error(mock.Anything, mock.Anything).Maybe()
			},
			expectedErr: fmt.Errorf("error creating user: %w", domain.ErrInvalidEmail),
		},
		{
			name: "user already exists",
			command: RegisterUserCommandV1{
				ID:       uuid.NewUuid().ToString(),
				Name:     "existinguser",
				Email:    "existing@example.com",
				Password: "password123",
			},
			setupMocks: func(mockRepo *userDomainPortsMocks.MockUserRepository, mockLogger *loggerMocks.MockLogger, mockPublisher *messagingMocks.MockUserMessagePublisher, mockHasher *userDomainPortsMocks.MockPasswordHasher) {
				mockHasher.EXPECT().HashPassword(mock.AnythingOfType("string")).Return("hashedpassword", nil).Once()
				mockRepo.EXPECT().IsExists(mock.Anything, mock.AnythingOfType("string")).Return(true, nil).Once()
				mockLogger.EXPECT().Info(mock.Anything, mock.Anything).Maybe()
			},
			expectedErr: domain.ErrAlreadyExists,
		},
		{
			name: "repository IsExists error",
			command: RegisterUserCommandV1{
				ID:       uuid.NewUuid().ToString(),
				Name:     "testuser",
				Email:    "test@example.com",
				Password: "password123",
			},
			setupMocks: func(mockRepo *userDomainPortsMocks.MockUserRepository, mockLogger *loggerMocks.MockLogger, mockPublisher *messagingMocks.MockUserMessagePublisher, mockHasher *userDomainPortsMocks.MockPasswordHasher) {
				mockHasher.EXPECT().HashPassword(mock.AnythingOfType("string")).Return("hashedpassword", nil).Once()
				mockRepo.EXPECT().IsExists(mock.Anything, mock.AnythingOfType("string")).Return(false, errors.New("db error")).Once()
				mockLogger.EXPECT().Error(mock.Anything, mock.Anything).Maybe()
			},
			expectedErr: fmt.Errorf("error checking if user exists: %w", errors.New("db error")),
		},
		{
			name: "repository Register error",
			command: RegisterUserCommandV1{
				ID:       uuid.NewUuid().ToString(),
				Name:     "testuser",
				Email:    "test@example.com",
				Password: "password123",
			},
			setupMocks: func(mockRepo *userDomainPortsMocks.MockUserRepository, mockLogger *loggerMocks.MockLogger, mockPublisher *messagingMocks.MockUserMessagePublisher, mockHasher *userDomainPortsMocks.MockPasswordHasher) {
				mockHasher.EXPECT().HashPassword(mock.AnythingOfType("string")).Return("hashedpassword", nil).Once()
				mockRepo.EXPECT().IsExists(mock.Anything, mock.AnythingOfType("string")).Return(false, nil).Once()
				mockRepo.EXPECT().Register(mock.Anything, mock.AnythingOfType("*domain.User")).Return(errors.New("repo save error")).Once()
				mockLogger.EXPECT().Error(mock.Anything, mock.Anything).Maybe()
			},
			expectedErr: errors.New("repo save error"),
		},
		{
			name: "message publishing error",
			command: RegisterUserCommandV1{
				ID:       uuid.NewUuid().ToString(),
				Name:     "testuser",
				Email:    "test@example.com",
				Password: "password123",
			},
			setupMocks: func(mockRepo *userDomainPortsMocks.MockUserRepository, mockLogger *loggerMocks.MockLogger, mockPublisher *messagingMocks.MockUserMessagePublisher, mockHasher *userDomainPortsMocks.MockPasswordHasher) {
				mockHasher.EXPECT().HashPassword(mock.AnythingOfType("string")).Return("hashedpassword", nil).Once()
				mockRepo.EXPECT().IsExists(mock.Anything, mock.AnythingOfType("string")).Return(false, nil).Once()
				mockRepo.EXPECT().Register(mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()
				mockPublisher.EXPECT().PublishUserRegistered(mock.Anything, mock.AnythingOfType("*domain.User")).Return(errors.New("publish error")).Once()
				mockLogger.EXPECT().Error(mock.Anything, mock.Anything).Maybe()
			},
			expectedErr: fmt.Errorf("failed to publish the user created event: %w", errors.New("publish error")),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Given
			mockRepo := userDomainPortsMocks.NewMockUserRepository(t)
			mockLogger := loggerMocks.NewMockLogger(t)
			mockPublisher := messagingMocks.NewMockUserMessagePublisher(t)
			mockHasher := userDomainPortsMocks.NewMockPasswordHasher(t)
			handler := NewRegisterUserHandler(mockRepo, mockLogger, mockPublisher, mockHasher)

			tc.setupMocks(mockRepo, mockLogger, mockPublisher, mockHasher)

			// When
			err := handler.Handle(context.Background(), tc.command)

			// Then
			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
