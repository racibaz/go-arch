package module

import (
	"testing"

	appMockPorts "github.com/racibaz/go-arch/internal/modules/post/testing/mocks/application/ports"
	domainMockPorts "github.com/racibaz/go-arch/internal/modules/post/testing/mocks/domain/ports"
	"github.com/racibaz/go-arch/pkg/logger"
	"github.com/stretchr/testify/suite"
)

// PostModuleTestSuite provides a test suite for PostModule testing
type PostModuleTestSuite struct {
	suite.Suite
	mockRepo    *domainMockPorts.MockPostRepository
	mockAdapter *domainMockPorts.MockNotificationAdapter
	mockService *appMockPorts.MockPostService
	mockLogger  *logger.MockLogger
}

// SetupTest is called before each test method
func (suite *PostModuleTestSuite) SetupTest() {
	suite.mockRepo = domainMockPorts.NewMockPostRepository(suite.T())
	suite.mockAdapter = domainMockPorts.NewMockNotificationAdapter(suite.T())
	suite.mockService = appMockPorts.NewMockPostService(suite.T())

	suite.mockLogger = logger.NewMockLogger(suite.T())
}

// TearDownTest is called after each test method
func (suite *PostModuleTestSuite) TearDownTest() {
	// Cleanup is handled by testify's mock cleanup
}

func TestPostModuleTestSuite(t *testing.T) {
	suite.Run(t, new(PostModuleTestSuite))
}

func (suite *PostModuleTestSuite) TestNewPostModule() {
	suite.Run("should create post module successfully", func() {
		// When
		postModule := NewPostModule(
			suite.mockRepo,
			suite.mockService,
			suite.mockLogger,
			suite.mockAdapter,
		)

		// Then
		suite.NotNil(postModule)
		suite.NotNil(postModule.Repository())
		suite.NotNil(postModule.Service())
		suite.NotNil(postModule.Notifier())
		// Logger can be nil for testing purposes
	})
}

func (suite *PostModuleTestSuite) TestPostModule_Repository() {
	suite.Run("should return configured repository", func() {
		// Given
		postModule := NewPostModule(
			suite.mockRepo,
			suite.mockService,
			suite.mockLogger,
			suite.mockAdapter,
		)

		// When
		repo := postModule.Repository()

		// Then
		suite.NotNil(repo)
		suite.Equal(suite.mockRepo, repo)
	})
}

func (suite *PostModuleTestSuite) TestPostModule_Service() {
	suite.Run("should return configured service", func() {
		// Given
		postModule := NewPostModule(
			suite.mockRepo,
			suite.mockService,
			suite.mockLogger,
			suite.mockAdapter,
		)

		// When
		service := postModule.Service()

		// Then
		suite.NotNil(service)
		suite.Equal(suite.mockService, service)
	})
}

func (suite *PostModuleTestSuite) TestPostModule_Logger() {
	suite.Run("should return configured logger", func() {
		// Given
		postModule := NewPostModule(
			suite.mockRepo,
			suite.mockService,
			suite.mockLogger,
			suite.mockAdapter,
		)

		// When
		logger := postModule.Logger()

		// Then
		suite.Equal(suite.mockLogger, logger)
	})
}

func (suite *PostModuleTestSuite) TestPostModule_Notifier() {
	suite.Run("should return configured notifier", func() {
		// Given
		postModule := NewPostModule(
			suite.mockRepo,
			suite.mockService,
			suite.mockLogger,
			suite.mockAdapter,
		)

		// When
		notifier := postModule.Notifier()

		// Then
		suite.NotNil(notifier)
		suite.Equal(suite.mockAdapter, notifier)
	})
}

func (suite *PostModuleTestSuite) TestPostModule_MultipleInstances() {
	suite.Run("should create multiple independent module instances", func() {
		// Given
		mockRepo2 := domainMockPorts.NewMockPostRepository(suite.T())
		mockService2 := appMockPorts.NewMockPostService(suite.T())
		mockAdapter2 := domainMockPorts.NewMockNotificationAdapter(suite.T())

		// When
		module1 := NewPostModule(
			suite.mockRepo,
			suite.mockService,
			suite.mockLogger,
			suite.mockAdapter,
		)
		module2 := NewPostModule(
			mockRepo2,
			mockService2,
			suite.mockLogger,
			mockAdapter2,
		)

		// Then
		suite.NotNil(module1)
		suite.NotNil(module2)
		suite.NotSame(module1, module2)
		suite.NotSame(module1.Repository(), module2.Repository())
		suite.NotSame(module1.Service(), module2.Service())
		suite.NotSame(module1.Notifier(), module2.Notifier())
	})
}

func (suite *PostModuleTestSuite) TestPostModule_WithNilDependencies() {
	suite.Run("should handle nil dependencies gracefully", func() {
		// When
		postModule := NewPostModule(nil, nil, nil, nil)

		// Then
		suite.NotNil(postModule)
		suite.Nil(postModule.Repository())
		suite.Nil(postModule.Service())
		suite.Nil(postModule.Logger())
		suite.Nil(postModule.Notifier())
	})
}

func (suite *PostModuleTestSuite) TestPostModule_DependencyInjection() {
	suite.Run("should properly inject all dependencies", func() {
		// When
		postModule := NewPostModule(
			suite.mockRepo,
			suite.mockService,
			suite.mockLogger,
			suite.mockAdapter,
		)

		// Then
		suite.Equal(suite.mockRepo, postModule.Repository())
		suite.Equal(suite.mockService, postModule.Service())
		suite.Equal(suite.mockLogger, postModule.Logger())
		suite.Equal(suite.mockAdapter, postModule.Notifier())
	})
}
