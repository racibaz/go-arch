package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	sharedPortsMocks "github.com/racibaz/go-arch/internal/modules/shared/testing/mocks/application/ports"
	userCommands "github.com/racibaz/go-arch/internal/modules/user/features/signup/v1/application/commands"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewRegisterUserHandler(t *testing.T) {
	// Given
	mockCommandHandler := sharedPortsMocks.NewMockCommandHandler[userCommands.RegisterUserCommandV1](t)

	// When
	handler := NewRegisterUserHandler(mockCommandHandler)

	// Then
	assert.NotNil(t, handler)
	assert.Equal(t, mockCommandHandler, handler.Handler)
}

func TestRegisterUserHandler_Store(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testCases := []struct {
		name             string
		requestBody      RegisterUserRequestDto
		expectedStatus   int
		setupMocks       func(*sharedPortsMocks.MockCommandHandler[userCommands.RegisterUserCommandV1])
		expectedResponse string // Can be a partial match or full match depending on complexity
	}{
		{
			name: "successful user registration",
			requestBody: RegisterUserRequestDto{
				Name:     "testuser",
				Email:    "test@example.com",
				Password: "password123",
			},
			expectedStatus: http.StatusCreated,
			setupMocks: func(m *sharedPortsMocks.MockCommandHandler[userCommands.RegisterUserCommandV1]) {
				m.EXPECT().Handle(mock.Anything, mock.AnythingOfType("commands.RegisterUserCommandV1")).Return(nil).Once()
			},
			expectedResponse: "User is registered successfully",
		},
		{
			name: "malformed json - decode error",
			expectedStatus: http.StatusBadRequest,
			setupMocks: func(m *sharedPortsMocks.MockCommandHandler[userCommands.RegisterUserCommandV1]) {
				// No handler call expected for decode error
			},
			expectedResponse: InValidErrMessage,
		},
		{
			name: "invalid request body - validation error",
			requestBody: RegisterUserRequestDto{
				Name:     "", // Invalid name
				Email:    "invalid-email",
				Password: "short",
			},
			expectedStatus: http.StatusBadRequest,
			setupMocks: func(m *sharedPortsMocks.MockCommandHandler[userCommands.RegisterUserCommandV1]) {
				// No handler call expected for validation error
			},
			expectedResponse: ValidationErrMessage,
		},
		{
			name: "handler returns error",
			requestBody: RegisterUserRequestDto{
				Name:     "testuser",
				Email:    "test@example.com",
				Password: "password123",
			},
			expectedStatus: http.StatusInternalServerError,
			setupMocks: func(m *sharedPortsMocks.MockCommandHandler[userCommands.RegisterUserCommandV1]) {
				m.EXPECT().Handle(mock.Anything, mock.AnythingOfType("commands.RegisterUserCommandV1")).Return(errors.New("internal server error")).Once()
			},
			expectedResponse: "user registration failed",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Given
			mockCommandHandler := sharedPortsMocks.NewMockCommandHandler[userCommands.RegisterUserCommandV1](t)
			setupMocksForRegisterUserHandler(mockCommandHandler, tc.setupMocks)

			receiver := NewRegisterUserHandler(mockCommandHandler)
			rr := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rr)

			// Marshal the request body
			jsonBody, err := json.Marshal(tc.requestBody)
			assert.NoError(t, err)

			// Create a new HTTP request
			var req *http.Request
			if tc.name == "malformed json - decode error" {
				req, err = http.NewRequest(http.MethodPost, routePath, bytes.NewBufferString("{invalid json"))
			} else {
				req, err = http.NewRequest(http.MethodPost, routePath, bytes.NewBuffer(jsonBody))
			}
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			// When
			receiver.Store(c)

			// Then
			assert.Equal(t, tc.expectedStatus, rr.Code)
			assert.Contains(t, rr.Body.String(), tc.expectedResponse)

			mockCommandHandler.AssertExpectations(t)
		})
	}
}

// Helper to setup mocks with correct generic type
func setupMocksForRegisterUserHandler(m *sharedPortsMocks.MockCommandHandler[userCommands.RegisterUserCommandV1], setup func(*sharedPortsMocks.MockCommandHandler[userCommands.RegisterUserCommandV1])) {
	if setup != nil {
		setup(m)
	}
}
