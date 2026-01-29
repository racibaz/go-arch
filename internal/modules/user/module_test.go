package module

import (
	"testing"

	sharedPortsMocks "github.com/racibaz/go-arch/internal/modules/shared/testing/mocks/application/ports"
	userCommands "github.com/racibaz/go-arch/internal/modules/user/features/signup/v1/application/commands"
	userQueries "github.com/racibaz/go-arch/internal/modules/user/features/login/v1/application/queries"
	userPortsMocks "github.com/racibaz/go-arch/internal/modules/user/testing/mocks/domain/ports"
	loggerMocks "github.com/racibaz/go-arch/pkg/logger"
	"github.com/racibaz/go-arch/pkg/ddd"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewUserModule(t *testing.T) {
	// Given
	mockRepo := userPortsMocks.NewMockUserRepository(t)
	mockSignupHandler := sharedPortsMocks.NewMockCommandHandler[userCommands.RegisterUserCommandV1](t)
	mockLoginHandler := sharedPortsMocks.NewMockQueryHandler[userQueries.LoginQueryV1, *userQueries.LoginQueryResponse](t)
	mockLogger := loggerMocks.NewMockLogger(t)
	mockNotifier := userPortsMocks.NewMockNotificationAdapter(t)

	// When
	userModule := NewUserModule(
		mockRepo,
		mockSignupHandler,
		mockLoginHandler,
		mockLogger,
		mockNotifier,
	)

	// Then
	assert.NotNil(t, userModule)
	assert.Equal(t, mockRepo, userModule.repository)
	assert.Equal(t, mockSignupHandler, userModule.signupCommandHandler)
	assert.Equal(t, mockLoginHandler, userModule.loginQueryHandler)
	assert.Equal(t, mockLogger, userModule.logger)
	assert.Equal(t, mockNotifier, userModule.notifier)
}

func TestUserModule_Accessors(t *testing.T) {
	// Given
	mockRepo := userPortsMocks.NewMockUserRepository(t)
	mockSignupHandler := sharedPortsMocks.NewMockCommandHandler[userCommands.RegisterUserCommandV1](t)
	mockLoginHandler := sharedPortsMocks.NewMockQueryHandler[userQueries.LoginQueryV1, *userQueries.LoginQueryResponse](t)
	mockLogger := loggerMocks.NewMockLogger(t)
	mockNotifier := userPortsMocks.NewMockNotificationAdapter(t)

	userModule := NewUserModule(
		mockRepo,
		mockSignupHandler,
		mockLoginHandler,
		mockLogger,
		mockNotifier,
	)

	// Then
	assert.Equal(t, mockRepo, userModule.Repository())
	assert.Equal(t, mockSignupHandler, userModule.RegisterUserCommandHandler())
	assert.Equal(t, mockLoginHandler, userModule.LoginQueryHandler())
	assert.Equal(t, mockLogger, userModule.Logger())
	assert.Equal(t, mockNotifier, userModule.Notifier())
}

func TestRegisterNotificationHandlers(t *testing.T) {
	// Given
	mockEventHandler := new(ddd.MockEventHandler[ddd.AggregateEvent])
	mockEventSubscriber := new(ddd.MockEventSubscriber[ddd.AggregateEvent])

	// Expected behavior: Subscribe should be called once with a specific event type and the handler
	mockEventSubscriber.EXPECT().Subscribe(mock.Anything, mock.Anything).Return().Once()

	// When
	RegisterNotificationHandlers(mockEventHandler, mockEventSubscriber)

	// Then
	mockEventSubscriber.AssertExpectations(t)
}
