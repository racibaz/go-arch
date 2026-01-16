package query

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/internal/modules/post/testing/mocks/domain/ports"
	applicationPorts "github.com/racibaz/go-arch/internal/modules/shared/application/ports"
	loggerMocks "github.com/racibaz/go-arch/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewGetPostHandler(t *testing.T) {
	// Given
	mockRepo := ports.NewMockPostRepository(t)
	mockLogger := loggerMocks.NewMockLogger(t)

	// When
	handler := NewGetPostHandler(mockRepo, mockLogger)

	// Then
	assert.NotNil(t, handler)
	assert.Equal(t, mockRepo, handler.PostRepository)
	assert.Equal(t, mockLogger, handler.logger)
	assert.NotNil(t, handler.tracer)
}

func TestGetPostHandler_ImplementsQueryHandlerInterface(t *testing.T) {
	// Given
	mockRepo := ports.NewMockPostRepository(t)
	mockLogger := loggerMocks.NewMockLogger(t)
	handler := NewGetPostHandler(mockRepo, mockLogger)

	// When & Then
	var _ applicationPorts.QueryHandler[GetPostByIdQuery, GetPostByIdQueryResponse] = handler
}

func TestGetPostHandler_Handle(t *testing.T) {
	testCases := []struct {
		name        string
		query       GetPostByIdQuery
		setupMocks  func(*ports.MockPostRepository, *loggerMocks.MockLogger)
		expected    GetPostByIdQueryResponse
		expectedErr error
	}{
		{
			name: "successful post retrieval",
			query: GetPostByIdQuery{
				ID: "test-post-id",
			},
			setupMocks: func(mockRepo *ports.MockPostRepository, mockLogger *loggerMocks.MockLogger) {
				createdAt := time.Now()
				updatedAt := time.Now()
				post := createTestPost("test-post-id", "test-user-id", createdAt, updatedAt)
				mockRepo.EXPECT().GetByID(mock.Anything, "test-post-id").Return(post, nil).Once()
				mockLogger.EXPECT().Info(mock.Anything, mock.Anything).Maybe()
			},
			expected: GetPostByIdQueryResponse{
				ID:          "test-post-id",
				UserID:      "test-user-id",
				Title:       "Test Title That Is Long Enough",
				Description: "Test Description That Is Long Enough",
				Content:     "Test Content That Is Long Enough",
				Status:      int(domain.PostStatusDraft),
			},
			expectedErr: nil,
		},
		{
			name: "post not found",
			query: GetPostByIdQuery{
				ID: "non-existent-id",
			},
			setupMocks: func(mockRepo *ports.MockPostRepository, mockLogger *loggerMocks.MockLogger) {
				mockRepo.EXPECT().
					GetByID(mock.Anything, "non-existent-id").
					Return(nil, domain.ErrNotFound).
					Once()
				mockLogger.EXPECT().Error(mock.Anything, mock.Anything).Maybe()
			},
			expected:    GetPostByIdQueryResponse{},
			expectedErr: domain.ErrNotFound,
		},
		{
			name: "repository error",
			query: GetPostByIdQuery{
				ID: "test-post-id",
			},
			setupMocks: func(mockRepo *ports.MockPostRepository, mockLogger *loggerMocks.MockLogger) {
				mockRepo.EXPECT().
					GetByID(mock.Anything, "test-post-id").
					Return(nil, errors.New("database connection error")).
					Once()
				mockLogger.EXPECT().Error(mock.Anything, mock.Anything).Maybe()
			},
			expected:    GetPostByIdQueryResponse{},
			expectedErr: errors.New("database connection error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Given
			mockRepo := ports.NewMockPostRepository(t)
			mockLogger := loggerMocks.NewMockLogger(t)
			handler := NewGetPostHandler(mockRepo, mockLogger)

			tc.setupMocks(mockRepo, mockLogger)

			// When
			result, err := handler.Handle(context.Background(), tc.query)

			// Then
			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErr.Error(), err.Error())
				assert.Equal(t, GetPostByIdQueryResponse{}, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected.ID, result.ID)
				assert.Equal(t, tc.expected.UserID, result.UserID)
				assert.Equal(t, tc.expected.Title, result.Title)
				assert.Equal(t, tc.expected.Description, result.Description)
				assert.Equal(t, tc.expected.Content, result.Content)
				assert.Equal(t, tc.expected.Status, result.Status)
				// Note: CreatedAt and UpdatedAt are set dynamically, so we don't compare them exactly
				assert.False(t, result.CreatedAt.IsZero())
				assert.False(t, result.UpdatedAt.IsZero())
			}
		})
	}
}

func TestGetPostHandler_Handle_ResponseMapping(t *testing.T) {
	// Test that all fields are properly mapped from domain to response
	createdAt := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
	updatedAt := time.Date(2023, 1, 2, 12, 0, 0, 0, time.UTC)

	// Given
	mockRepo := ports.NewMockPostRepository(t)
	mockLogger := loggerMocks.NewMockLogger(t)
	handler := NewGetPostHandler(mockRepo, mockLogger)

	post := createTestPost("test-id", "user-id", createdAt, updatedAt)
	mockRepo.EXPECT().GetByID(mock.Anything, "test-id").Return(post, nil).Once()

	// When
	result, err := handler.Handle(context.Background(), GetPostByIdQuery{ID: "test-id"})

	// Then
	assert.NoError(t, err)
	assert.Equal(t, "test-id", result.ID)
	assert.Equal(t, "user-id", result.UserID)
	assert.Equal(t, "Test Title That Is Long Enough", result.Title)
	assert.Equal(t, "Test Description That Is Long Enough", result.Description)
	assert.Equal(t, "Test Content That Is Long Enough", result.Content)
	assert.Equal(t, int(domain.PostStatusDraft), result.Status)
	assert.Equal(t, createdAt, result.CreatedAt)
	assert.Equal(t, updatedAt, result.UpdatedAt)
}

func TestGetPostHandler_Handle_Tracing(t *testing.T) {
	// Given
	mockRepo := ports.NewMockPostRepository(t)
	mockLogger := loggerMocks.NewMockLogger(t)
	handler := NewGetPostHandler(mockRepo, mockLogger)

	createdAt := time.Now()
	updatedAt := time.Now()
	post := createTestPost("test-id", "user-id", createdAt, updatedAt)
	mockRepo.EXPECT().GetByID(mock.Anything, "test-id").Return(post, nil).Once()

	// When
	ctx := context.Background()
	result, err := handler.Handle(ctx, GetPostByIdQuery{ID: "test-id"})

	// Then
	assert.NoError(t, err)
	assert.NotEmpty(t, result.ID)
	// Note: Tracing verification would require more complex setup with OpenTelemetry
	// For now, we just ensure the handler completes without error
}

// Helper function to create a test post
func createTestPost(id, userID string, createdAt, updatedAt time.Time) *domain.Post {
	post, _ := domain.Create(
		id,
		userID,
		"Test Title That Is Long Enough",
		"Test Description That Is Long Enough",
		"Test Content That Is Long Enough",
		domain.PostStatusDraft,
		createdAt,
		updatedAt,
	)
	return post
}
