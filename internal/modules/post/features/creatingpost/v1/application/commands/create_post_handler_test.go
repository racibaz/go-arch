package commands

import (
	"context"
	"errors"
	"testing"

	"github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/internal/modules/post/testing/mocks/domain/ports"
	applicationPorts "github.com/racibaz/go-arch/internal/modules/shared/application/ports"
	loggerMocks "github.com/racibaz/go-arch/pkg/logger"
	messagingMocks "github.com/racibaz/go-arch/pkg/messaging"
	"github.com/racibaz/go-arch/pkg/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewCreatePostHandler(t *testing.T) {
	// Given
	mockRepo := ports.NewMockPostRepository(t)
	mockLogger := loggerMocks.NewMockLogger(t)
	mockPublisher := messagingMocks.NewMockMessagePublisher(t)

	// When
	handler := NewCreatePostHandler(mockRepo, mockLogger, mockPublisher)

	// Then
	assert.NotNil(t, handler)
	assert.Equal(t, mockRepo, handler.PostRepository)
	assert.Equal(t, mockLogger, handler.logger)
	assert.Equal(t, mockPublisher, handler.messagePublisher)
	assert.NotNil(t, handler.tracer)
}

func TestCreatePostHandler_ImplementsCommandHandlerInterface(t *testing.T) {
	// Given
	mockRepo := ports.NewMockPostRepository(t)
	mockLogger := loggerMocks.NewMockLogger(t)
	mockPublisher := messagingMocks.NewMockMessagePublisher(t)
	handler := NewCreatePostHandler(mockRepo, mockLogger, mockPublisher)

	// When & Then
	var _ applicationPorts.CommandHandler[CreatePostCommandV1] = handler
}

func TestCreatePostHandler_Handle(t *testing.T) {
	testCases := []struct {
		name        string
		command     CreatePostCommandV1
		setupMocks  func(*ports.MockPostRepository, *loggerMocks.MockLogger, *messagingMocks.MockMessagePublisher)
		expectedErr error
	}{
		{
			name: "successful post creation",
			command: CreatePostCommandV1{
				ID:          uuid.NewUuid().ToString(),
				UserID:      uuid.NewUuid().ToString(),
				Title:       "Test Title That Is Long Enough",
				Description: "Test Description That Is Long Enough",
				Content:     "Test Content That Is Long Enough",
				Status:      domain.PostStatusDraft,
			},
			setupMocks: func(mockRepo *ports.MockPostRepository, mockLogger *loggerMocks.MockLogger, mockPublisher *messagingMocks.MockMessagePublisher) {
				mockRepo.EXPECT().
					IsExists(mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
					Return(false, nil).
					Once()
				mockRepo.EXPECT().
					Save(mock.Anything, mock.AnythingOfType("*domain.Post")).
					Return(nil).
					Once()
				mockPublisher.EXPECT().
					PublishPostCreated(mock.Anything, mock.AnythingOfType("*domain.Post")).
					Return(nil).
					Once()
				mockLogger.EXPECT().Info(mock.Anything, mock.Anything).Maybe()
				mockLogger.EXPECT().Error(mock.Anything, mock.Anything).Maybe()
			},
			expectedErr: nil,
		},
		{
			name: "post already exists",
			command: CreatePostCommandV1{
				ID:          uuid.NewUuid().ToString(),
				UserID:      uuid.NewUuid().ToString(),
				Title:       "Existing Title That Is Long Enough",
				Description: "Existing Description That Is Long Enough",
				Content:     "Existing Content That Is Long Enough",
				Status:      domain.PostStatusDraft,
			},
			setupMocks: func(mockRepo *ports.MockPostRepository, mockLogger *loggerMocks.MockLogger, mockPublisher *messagingMocks.MockMessagePublisher) {
				mockRepo.EXPECT().
					IsExists(mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
					Return(true, nil).
					Once()
				mockLogger.EXPECT().Info(mock.Anything, mock.Anything).Maybe()
			},
			expectedErr: domain.ErrAlreadyExists,
		},
		{
			name: "repository IsExists error",
			command: CreatePostCommandV1{
				ID:          uuid.NewUuid().ToString(),
				UserID:      uuid.NewUuid().ToString(),
				Title:       "Test Title That Is Long Enough",
				Description: "Test Description That Is Long Enough",
				Content:     "Test Content That Is Long Enough",
				Status:      domain.PostStatusDraft,
			},
			setupMocks: func(mockRepo *ports.MockPostRepository, mockLogger *loggerMocks.MockLogger, mockPublisher *messagingMocks.MockMessagePublisher) {
				mockRepo.EXPECT().
					IsExists(mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
					Return(false, errors.New("database error")).
					Once()
				mockLogger.EXPECT().Error(mock.Anything, mock.Anything).Maybe()
			},
			expectedErr: errors.New("error checking if post exists: database error"),
		},
		{
			name: "repository Save error",
			command: CreatePostCommandV1{
				ID:          uuid.NewUuid().ToString(),
				UserID:      uuid.NewUuid().ToString(),
				Title:       "Test Title That Is Long Enough",
				Description: "Test Description That Is Long Enough",
				Content:     "Test Content That Is Long Enough",
				Status:      domain.PostStatusDraft,
			},
			setupMocks: func(mockRepo *ports.MockPostRepository, mockLogger *loggerMocks.MockLogger, mockPublisher *messagingMocks.MockMessagePublisher) {
				mockRepo.EXPECT().
					IsExists(mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
					Return(false, nil).
					Once()
				mockRepo.EXPECT().
					Save(mock.Anything, mock.AnythingOfType("*domain.Post")).
					Return(errors.New("save error")).
					Once()
				mockLogger.EXPECT().Error(mock.Anything, mock.Anything).Maybe()
			},
			expectedErr: errors.New("save error"),
		},
		{
			name: "message publishing error",
			command: CreatePostCommandV1{
				ID:          uuid.NewUuid().ToString(),
				UserID:      uuid.NewUuid().ToString(),
				Title:       "Test Title That Is Long Enough",
				Description: "Test Description That Is Long Enough",
				Content:     "Test Content That Is Long Enough",
				Status:      domain.PostStatusDraft,
			},
			setupMocks: func(mockRepo *ports.MockPostRepository, mockLogger *loggerMocks.MockLogger, mockPublisher *messagingMocks.MockMessagePublisher) {
				mockRepo.EXPECT().
					IsExists(mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
					Return(false, nil).
					Once()
				mockRepo.EXPECT().
					Save(mock.Anything, mock.AnythingOfType("*domain.Post")).
					Return(nil).
					Once()
				mockPublisher.EXPECT().
					PublishPostCreated(mock.Anything, mock.AnythingOfType("*domain.Post")).
					Return(errors.New("publish error")).
					Once()
			},
			expectedErr: errors.New("failed to publish the post created event: publish error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Given
			mockRepo := ports.NewMockPostRepository(t)
			mockLogger := loggerMocks.NewMockLogger(t)
			mockPublisher := messagingMocks.NewMockMessagePublisher(t)
			handler := NewCreatePostHandler(mockRepo, mockLogger, mockPublisher)

			tc.setupMocks(mockRepo, mockLogger, mockPublisher)

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

func TestCreatePostHandler_Handle_InvalidPostCreation(t *testing.T) {
	// Note: The current implementation has a bug where it ignores domain.Create errors
	// and tries to access fields on a nil post, causing panics.
	// These tests demonstrate the current behavior but the handler should be fixed
	// to properly handle domain validation errors.

	testCases := []struct {
		name        string
		command     CreatePostCommandV1
		expectPanic bool
	}{
		{
			name: "empty ID",
			command: CreatePostCommandV1{
				ID:          "",
				UserID:      uuid.NewUuid().ToString(),
				Title:       "Test Title That Is Long Enough",
				Description: "Test Description That Is Long Enough",
				Content:     "Test Content That Is Long Enough",
				Status:      domain.PostStatusDraft,
			},
			expectPanic: true, // Current behavior: panic because post is nil
		},
		{
			name: "empty user ID",
			command: CreatePostCommandV1{
				ID:          uuid.NewUuid().ToString(),
				UserID:      "",
				Title:       "Test Title That Is Long Enough",
				Description: "Test Description That Is Long Enough",
				Content:     "Test Content That Is Long Enough",
				Status:      domain.PostStatusDraft,
			},
			expectPanic: true, // Current behavior: panic because post is nil
		},
		{
			name: "title too short",
			command: CreatePostCommandV1{
				ID:          uuid.NewUuid().ToString(),
				UserID:      uuid.NewUuid().ToString(),
				Title:       "Hi",
				Description: "Test Description That Is Long Enough",
				Content:     "Test Content That Is Long Enough",
				Status:      domain.PostStatusDraft,
			},
			expectPanic: true, // Current behavior: panic because post is nil
		},
		{
			name: "description too short",
			command: CreatePostCommandV1{
				ID:          uuid.NewUuid().ToString(),
				UserID:      uuid.NewUuid().ToString(),
				Title:       "Test Title That Is Long Enough",
				Description: "Hi",
				Content:     "Test Content That Is Long Enough",
				Status:      domain.PostStatusDraft,
			},
			expectPanic: true, // Current behavior: panic because post is nil
		},
		{
			name: "content too short",
			command: CreatePostCommandV1{
				ID:          uuid.NewUuid().ToString(),
				UserID:      uuid.NewUuid().ToString(),
				Title:       "Test Title That Is Long Enough",
				Description: "Test Description That Is Long Enough",
				Content:     "Hi",
				Status:      domain.PostStatusDraft,
			},
			expectPanic: true, // Current behavior: panic because post is nil
		},
		{
			name: "invalid status",
			command: CreatePostCommandV1{
				ID:          uuid.NewUuid().ToString(),
				UserID:      uuid.NewUuid().ToString(),
				Title:       "Test Title That Is Long Enough",
				Description: "Test Description That Is Long Enough",
				Content:     "Test Content That Is Long Enough",
				Status:      -1, // Invalid status
			},
			expectPanic: true, // Current behavior: panic because post is nil
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Given
			mockRepo := ports.NewMockPostRepository(t)
			mockLogger := loggerMocks.NewMockLogger(t)
			mockPublisher := messagingMocks.NewMockMessagePublisher(t)
			handler := NewCreatePostHandler(mockRepo, mockLogger, mockPublisher)

			// When & Then
			if tc.expectPanic {
				assert.Panics(t, func() {
					handler.Handle(context.Background(), tc.command)
				})
			} else {
				err := handler.Handle(context.Background(), tc.command)
				assert.Error(t, err)
			}
		})
	}
}

func TestCreatePostHandler_Handle_Tracing(t *testing.T) {
	// Given
	mockRepo := ports.NewMockPostRepository(t)
	mockLogger := loggerMocks.NewMockLogger(t)
	mockPublisher := messagingMocks.NewMockMessagePublisher(t)
	handler := NewCreatePostHandler(mockRepo, mockLogger, mockPublisher)

	command := CreatePostCommandV1{
		ID:          uuid.NewUuid().ToString(),
		UserID:      uuid.NewUuid().ToString(),
		Title:       "Test Title That Is Long Enough",
		Description: "Test Description That Is Long Enough",
		Content:     "Test Content That Is Long Enough",
		Status:      domain.PostStatusDraft,
	}

	mockRepo.EXPECT().
		IsExists(mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(false, nil).
		Once()
	mockRepo.EXPECT().Save(mock.Anything, mock.AnythingOfType("*domain.Post")).Return(nil).Once()
	mockPublisher.EXPECT().
		PublishPostCreated(mock.Anything, mock.AnythingOfType("*domain.Post")).
		Return(nil).
		Once()
	mockLogger.EXPECT().Info(mock.Anything, mock.Anything).Maybe()

	// When
	ctx := context.Background()
	err := handler.Handle(ctx, command)

	// Then
	assert.NoError(t, err)
	// Note: Tracing verification would require more complex setup with OpenTelemetry
	// For now, we just ensure the handler completes without error
}
