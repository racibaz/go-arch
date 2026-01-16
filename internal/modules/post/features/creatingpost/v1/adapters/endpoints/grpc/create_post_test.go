package grpc

import (
	"context"
	"errors"
	"testing"

	"github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/internal/modules/post/features/creatingpost/v1/adapters/endpoints/grpc/proto"
	"github.com/racibaz/go-arch/internal/modules/post/features/creatingpost/v1/application/commands"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

// MockCommandHandler is a mock implementation of CommandHandler for testing
type MockCommandHandler struct {
	mock.Mock
}

func (m *MockCommandHandler) Handle(ctx context.Context, cmd commands.CreatePostCommandV1) error {
	args := m.Called(ctx, cmd)
	return args.Error(0)
}

func TestNewCreatePostHandler(t *testing.T) {
	// Given
	mockHandler := &MockCommandHandler{}

	// When
	server := grpc.NewServer()
	NewCreatePostHandler(server, mockHandler)

	// Then
	// The constructor registers the handler with the gRPC server
	// We can't easily test the internal registration, but we can verify no panics occur
	// This test mainly ensures the constructor doesn't panic
	assert.NotNil(t, server)
}

func TestCreatePostHandler_CreatePost_Success(t *testing.T) {
	// Given
	mockHandler := &MockCommandHandler{}
	handler := &CreatePostHandler{
		Handler: mockHandler,
	}

	// Mock the handler to return success
	mockHandler.On("Handle", mock.Anything, mock.AnythingOfType("commands.CreatePostCommandV1")).
		Return(nil).
		Once()

	// When
	req := &proto.CreatePostInput{
		UserID:      "550e8400-e29b-41d4-a716-446655440000",
		Title:       "Test Title That Is Long Enough",
		Description: "Test Description That Is Long Enough",
		Content:     "Test Content That Is Long Enough",
	}

	resp, err := handler.CreatePost(context.Background(), req)

	// Then
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Id)

	mockHandler.AssertExpectations(t)
}

func TestCreatePostHandler_CreatePost_HandlerError(t *testing.T) {
	// Given
	mockHandler := &MockCommandHandler{}
	handler := &CreatePostHandler{
		Handler: mockHandler,
	}

	// Mock the handler to return an error
	expectedError := errors.New("handler error")
	mockHandler.On("Handle", mock.Anything, mock.AnythingOfType("commands.CreatePostCommandV1")).
		Return(expectedError).
		Once()

	// When
	req := &proto.CreatePostInput{
		UserID:      "550e8400-e29b-41d4-a716-446655440000",
		Title:       "Test Title That Is Long Enough",
		Description: "Test Description That Is Long Enough",
		Content:     "Test Content That Is Long Enough",
	}

	resp, err := handler.CreatePost(context.Background(), req)

	// Then
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Nil(t, resp)

	mockHandler.AssertExpectations(t)
}

func TestCreatePostHandler_CreatePost_CommandStructure(t *testing.T) {
	// Test that the correct command is constructed from the gRPC request
	mockHandler := &MockCommandHandler{}
	handler := &CreatePostHandler{
		Handler: mockHandler,
	}

	var capturedCommand commands.CreatePostCommandV1
	mockHandler.On("Handle", mock.Anything, mock.MatchedBy(func(cmd commands.CreatePostCommandV1) bool {
		capturedCommand = cmd
		return true
	})).
		Return(nil).
		Once()

	// When
	req := &proto.CreatePostInput{
		UserID:      "550e8400-e29b-41d4-a716-446655440000",
		Title:       "Test Title That Is Long Enough",
		Description: "Test Description That Is Long Enough",
		Content:     "Test Content That Is Long Enough",
	}

	resp, err := handler.CreatePost(context.Background(), req)

	// Then
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	// Verify command structure
	assert.NotEmpty(t, capturedCommand.ID)
	assert.Equal(t, req.UserID, capturedCommand.UserID)
	assert.Equal(t, req.Title, capturedCommand.Title)
	assert.Equal(t, req.Description, capturedCommand.Description)
	assert.Equal(t, req.Content, capturedCommand.Content)
	assert.Equal(t, domain.PostStatusDraft, capturedCommand.Status)
	assert.True(t, capturedCommand.CreatedAt.IsZero()) // gRPC handler sets zero time
	assert.True(t, capturedCommand.UpdatedAt.IsZero()) // gRPC handler sets zero time

	mockHandler.AssertExpectations(t)
}

func TestCreatePostHandler_CreatePost_ResponseStructure(t *testing.T) {
	// Given
	mockHandler := &MockCommandHandler{}
	handler := &CreatePostHandler{
		Handler: mockHandler,
	}

	mockHandler.On("Handle", mock.Anything, mock.AnythingOfType("commands.CreatePostCommandV1")).
		Return(nil).
		Once()

	// When
	req := &proto.CreatePostInput{
		UserID:      "550e8400-e29b-41d4-a716-446655440000",
		Title:       "Test Title That Is Long Enough",
		Description: "Test Description That Is Long Enough",
		Content:     "Test Content That Is Long Enough",
	}

	resp, err := handler.CreatePost(context.Background(), req)

	// Then
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Id)

	// Verify response is a valid UUID (since we use uuid.NewID())
	assert.Regexp(
		t,
		`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`,
		resp.Id,
	)

	mockHandler.AssertExpectations(t)
}

func TestCreatePostHandler_CreatePost_WithEmptyFields(t *testing.T) {
	// Test handling of empty fields in the request
	mockHandler := &MockCommandHandler{}
	handler := &CreatePostHandler{
		Handler: mockHandler,
	}

	mockHandler.On("Handle", mock.Anything, mock.AnythingOfType("commands.CreatePostCommandV1")).
		Return(nil).
		Once()

	// When - request with empty fields
	req := &proto.CreatePostInput{
		UserID:      "",
		Title:       "",
		Description: "",
		Content:     "",
	}

	resp, err := handler.CreatePost(context.Background(), req)

	// Then
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Id)

	mockHandler.AssertExpectations(t)
}

func TestCreatePostHandler_CreatePost_Tracing(t *testing.T) {
	// Given
	mockHandler := &MockCommandHandler{}
	handler := &CreatePostHandler{
		Handler: mockHandler,
	}

	mockHandler.On("Handle", mock.Anything, mock.AnythingOfType("commands.CreatePostCommandV1")).
		Return(nil).
		Once()

	// When
	ctx := context.Background()
	req := &proto.CreatePostInput{
		UserID:      "550e8400-e29b-41d4-a716-446655440000",
		Title:       "Test Title That Is Long Enough",
		Description: "Test Description That Is Long Enough",
		Content:     "Test Content That Is Long Enough",
	}

	resp, err := handler.CreatePost(ctx, req)

	// Then
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	// Tracing is tested implicitly - if tracing setup was broken, the handler would panic
	mockHandler.AssertExpectations(t)
}

func TestCreatePostHandler_CreatePost_ImplementsGRPCService(t *testing.T) {
	// Test that our handler properly implements the gRPC service interface
	mockHandler := &MockCommandHandler{}
	handler := &CreatePostHandler{
		Handler: mockHandler,
	}

	// Verify it implements the PostServiceServer interface
	var _ proto.PostServiceServer = handler
}

func TestCreatePostHandler_UnimplementedMethods(t *testing.T) {
	// Test that unimplemented methods return proper gRPC errors
	mockHandler := &MockCommandHandler{}
	handler := &CreatePostHandler{
		Handler: mockHandler,
	}

	// The handler embeds UnimplementedPostServiceServer which should handle
	// any methods not explicitly implemented
	// This is more of a compile-time check, but we can verify the struct composition
	assert.NotNil(t, handler.UnimplementedPostServiceServer)
}
