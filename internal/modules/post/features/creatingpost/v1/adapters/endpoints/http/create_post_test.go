package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/internal/modules/post/features/creatingpost/v1/application/commands"
	"github.com/racibaz/go-arch/pkg/helper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCommandHandler is a mock implementation of CreatePostCommandHandler for testing
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
	handler := NewCreatePostHandler(mockHandler)

	// Then
	assert.NotNil(t, handler)
	assert.Equal(t, mockHandler, handler.Handler)
}

func TestCreatePostHandler_Store_Success(t *testing.T) {
	var ( // Declare variables once at the top of the test case
		response          map[string]interface{}
		err               error
		responseDataBytes []byte
		apiResponse       helper.Response[CreatePostResponseDto]
	)

	// Given
	mockHandler := &MockCommandHandler{}
	handler := NewCreatePostHandler(mockHandler)

	// Mock the handler to return success
	mockHandler.On("Handle", mock.Anything, mock.AnythingOfType("commands.CreatePostCommandV1")).
		Return(nil).
		Once()

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/v1/posts", handler.Store)

	requestBody := CreatePostRequestDto{
		UserId:      "550e8400-e29b-41d4-a716-446655440000",
		Title:       "Test Title That Is Long Enough",
		Description: "Test Description That Is Long Enough",
		Content:     "Test Content That Is Long Enough",
	}

	jsonBody, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/api/v1/posts", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	// When
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Then
	assert.Equal(t, http.StatusCreated, w.Code)

	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, float64(http.StatusCreated), response["status"])
	assert.Equal(t, "Post created successfully", response["message"])

	// Now unmarshal the 'data' part into the specific DTO
	responseDataBytes, err = json.Marshal(response["data"])
	assert.NoError(t, err)

	err = json.Unmarshal(responseDataBytes, &apiResponse)
	assert.NoError(t, err)

	// Verify data structure
	assert.NotNil(t, apiResponse.Data)
	assert.NotNil(t, apiResponse.Data.Post)
	assert.Equal(t, requestBody.Title, apiResponse.Data.Post.Title)
	assert.Equal(t, requestBody.Description, apiResponse.Data.Post.Description)
	assert.Equal(t, requestBody.Content, apiResponse.Data.Post.Content)
	assert.Equal(t, domain.PostStatusDraft.String(), apiResponse.Data.Post.Status)

	// Verify HATEOAS links
	assert.NotNil(t, apiResponse.Links)
	assert.Equal(t, 4, len(apiResponse.Links)) // self, store, update, delete

	// Verify self link contains the newly created post ID
	var selfLinkID string
	for _, link := range apiResponse.Links {
		if link.Rel == "self" {
			selfLinkID = link.Href[strings.LastIndex(link.Href, "/")+1:]
			break
		}
	}
	assert.NotEmpty(t, selfLinkID)

	mockHandler.AssertExpectations(t)
}

func TestCreatePostHandler_Store_InvalidJSON(t *testing.T) {
	// Given
	mockHandler := &MockCommandHandler{}
	handler := NewCreatePostHandler(mockHandler)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/v1/posts", handler.Store)

	// Invalid JSON
	req, _ := http.NewRequest("POST", "/api/v1/posts", bytes.NewBufferString("invalid json"))
	req.Header.Set("Content-Type", "application/json")

	// When
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Then
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Invalid request body", response["message"])
	assert.Equal(t, float64(http.StatusBadRequest), response["status"])
}

func TestCreatePostHandler_Store_ValidationError(t *testing.T) {
	// NOTE: This test documents a bug in the current implementation.
	// The handler calls helper.ValidationErrorResponse but doesn't return,
	// so execution continues and the command handler is still called.
	// Due to this bug, the response contains malformed JSON from multiple responses.

	// Given
	mockHandler := &MockCommandHandler{}
	handler := NewCreatePostHandler(mockHandler)

	// No handler call expected for validation error

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/v1/posts", handler.Store)

	// Invalid request body (missing required fields)
	requestBody := CreatePostRequestDto{
		UserId: "550e8400-e29b-41d4-a716-446655440000",
		Title:  "Short", // Too short
	}

	jsonBody, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/api/v1/posts", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	// When
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Then
	// The validation error response is sent first (400)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, float64(http.StatusBadRequest), response["status"])
	assert.Equal(t, ValidationErrMessage, response["message"])
	assert.Contains(t, response["cause"].(map[string]interface{})["Title"].([]interface{})[0].(string), "min")

	mockHandler.AssertNotCalled(t, "Handle", mock.Anything, mock.Anything)
}

func TestCreatePostHandler_Store_HandlerError(t *testing.T) {
	// Given
	mockHandler := &MockCommandHandler{}
	handler := NewCreatePostHandler(mockHandler)

	// Mock the handler to return an error
	mockHandler.On("Handle", mock.Anything, mock.AnythingOfType("commands.CreatePostCommandV1")).
		Return(errors.New("handler error")).
		Once()

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/v1/posts", handler.Store)

	requestBody := CreatePostRequestDto{
		UserId:      "550e8400-e29b-41d4-a716-446655440000",
		Title:       "Test Title That Is Long Enough",
		Description: "Test Description That Is Long Enough",
		Content:     "Test Content That Is Long Enough",
	}

	jsonBody, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/api/v1/posts", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	// When
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Then
	// The handler error causes a 500 response
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	// Due to the response handling in Gin, we mainly verify that the handler was called
	mockHandler.AssertExpectations(t)
}

func TestCreatePostHandler_Store_CommandStructure(t *testing.T) {
	// Test that the correct command is passed to the handler
	mockHandler := &MockCommandHandler{}
	handler := NewCreatePostHandler(mockHandler)

	var capturedCommand commands.CreatePostCommandV1
	mockHandler.On("Handle", mock.Anything, mock.MatchedBy(func(cmd commands.CreatePostCommandV1) bool {
		capturedCommand = cmd
		return true
	})).
		Return(nil).
		Once()

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/v1/posts", handler.Store)

	requestBody := CreatePostRequestDto{
		UserId:      "550e8400-e29b-41d4-a716-446655440000",
		Title:       "Test Title That Is Long Enough",
		Description: "Test Description That Is Long Enough",
		Content:     "Test Content That Is Long Enough",
	}

	jsonBody, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/api/v1/posts", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	// When
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Then
	assert.Equal(t, http.StatusCreated, w.Code)

	// Verify command structure
	assert.NotEmpty(t, capturedCommand.ID)
	assert.Equal(t, requestBody.UserId, capturedCommand.UserID)
	assert.Equal(t, requestBody.Title, capturedCommand.Title)
	assert.Equal(t, requestBody.Description, capturedCommand.Description)
	assert.Equal(t, requestBody.Content, capturedCommand.Content)
	assert.Equal(t, domain.PostStatusDraft, capturedCommand.Status)
	assert.False(t, capturedCommand.CreatedAt.IsZero())
	assert.False(t, capturedCommand.UpdatedAt.IsZero())

	mockHandler.AssertExpectations(t)
}

func TestCreatePostHandler_Store_HATEOASLinks(t *testing.T) {
	// Given
	mockHandler := &MockCommandHandler{}
	handler := NewCreatePostHandler(mockHandler)

	mockHandler.On("Handle", mock.Anything, mock.AnythingOfType("commands.CreatePostCommandV1")).
		Return(nil).
		Once()

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/v1/posts", handler.Store)

	requestBody := CreatePostRequestDto{
		UserId:      "550e8400-e29b-41d4-a716-446655440000",
		Title:       "Test Title That Is Long Enough",
		Description: "Test Description That Is Long Enough",
		Content:     "Test Content That Is Long Enough",
	}

	jsonBody, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/api/v1/posts", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	// When
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Then
	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	responseData, ok := response["data"].(map[string]interface{})
	assert.True(t, ok)

	links, ok := responseData["_links"].([]interface{})
	assert.True(t, ok)
	assert.Equal(t, 4, len(links))

	// Verify each link
	linkMap := make(map[string]map[string]interface{})
	for _, linkInterface := range links {
		link := linkInterface.(map[string]interface{})
		rel := link["rel"].(string)
		linkMap[rel] = link
	}

	// Check self link
	selfLink, exists := linkMap["self"]
	assert.True(t, exists)
	assert.Contains(t, selfLink["href"].(string), "/api/v1/posts/")
	assert.Equal(t, "GET", selfLink["type"])

	// Check store link
	storeLink, exists := linkMap["store"]
	assert.True(t, exists)
	assert.Equal(t, "/api/v1/posts/", storeLink["href"])
	assert.Equal(t, "POST", storeLink["type"])

	// Check update link
	updateLink, exists := linkMap["update"]
	assert.True(t, exists)
	assert.Contains(t, updateLink["href"].(string), "/api/v1/posts/")
	assert.Equal(t, "PUT", updateLink["type"])

	// Check delete link
	deleteLink, exists := linkMap["delete"]
	assert.True(t, exists)
	assert.Contains(t, deleteLink["href"].(string), "/api/v1/posts/")
	assert.Equal(t, "DELETE", deleteLink["type"])

	mockHandler.AssertExpectations(t)
}

func TestCreatePostHandler_Store_ResponseStructure(t *testing.T) {
	// Given
	mockHandler := &MockCommandHandler{}
	handler := NewCreatePostHandler(mockHandler)

	mockHandler.On("Handle", mock.Anything, mock.AnythingOfType("commands.CreatePostCommandV1")).
		Return(nil).
		Once()

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/v1/posts", handler.Store)

	requestBody := CreatePostRequestDto{
		UserId:      "550e8400-e29b-41d4-a716-446655440000",
		Title:       "Test Title That Is Long Enough",
		Description: "Test Description That Is Long Enough",
		Content:     "Test Content That Is Long Enough",
	}

	jsonBody, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/api/v1/posts", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	// When
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Then
	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Verify top-level structure
	assert.Contains(t, response, "status")
	assert.Contains(t, response, "message")
	assert.Contains(t, response, "data")
	assert.Equal(t, float64(http.StatusCreated), response["status"])
	assert.Equal(t, "Post created successfully", response["message"])

	// Now unmarshal the 'data' part into the specific DTO
	responseDataBytes, err := json.Marshal(response["data"])
	assert.NoError(t, err)

	var apiResponse helper.Response[CreatePostResponseDto]
	err = json.Unmarshal(responseDataBytes, &apiResponse)
	assert.NoError(t, err)

	// Verify data structure
	assert.NotNil(t, apiResponse.Data)
	assert.NotNil(t, apiResponse.Data.Post)
	assert.Contains(t, apiResponse.Data.Post.Title, requestBody.Title)
	assert.Contains(t, apiResponse.Data.Post.Description, requestBody.Description)
	assert.Contains(t, apiResponse.Data.Post.Content, requestBody.Content)
	assert.Contains(t, apiResponse.Data.Post.Status, domain.PostStatusDraft.String())

	assert.NotNil(t, apiResponse.Links)
	mockHandler.AssertExpectations(t)
}