package module

import (
	"testing"

	"github.com/racibaz/go-arch/internal/modules/post/features/creatingpost/v1/application/commands"
	"github.com/racibaz/go-arch/internal/modules/post/features/gettingpostbyid/v1/application/query"
	gettingpostsQuery "github.com/racibaz/go-arch/internal/modules/post/features/gettingposts/v1/application/query"
	domainMockPorts "github.com/racibaz/go-arch/internal/modules/post/testing/mocks/domain/ports"
	appMockPorts "github.com/racibaz/go-arch/internal/modules/shared/testing/mocks/application/ports"
	"github.com/racibaz/go-arch/pkg/helper"
	"github.com/racibaz/go-arch/pkg/logger"
	"github.com/stretchr/testify/suite"
)

// PostModuleTestSuite provides a test suite for PostModule testing
type PostModuleTestSuite struct {
	suite.Suite
	mockRepo             *domainMockPorts.MockPostRepository
	mockAdapter          *domainMockPorts.MockNotificationAdapter
	mockCommandHandler   *appMockPorts.MockCommandHandler[commands.CreatePostCommandV1]
	mockQueryHandler     *appMockPorts.MockQueryHandler[query.GetPostByIdQuery, query.GetPostByIdQueryResponse]
	mockListQueryHandler *appMockPorts.MockQueryHandler[helper.Pagination, gettingpostsQuery.GetPostsQueryResponse]
	mockLogger           *logger.MockLogger
}

// SetupTest is called before each test method
func (suite *PostModuleTestSuite) SetupTest() {
	suite.mockRepo = domainMockPorts.NewMockPostRepository(suite.T())
	suite.mockAdapter = domainMockPorts.NewMockNotificationAdapter(suite.T())
	suite.mockCommandHandler = appMockPorts.NewMockCommandHandler[commands.CreatePostCommandV1](
		suite.T(),
	)
	suite.mockQueryHandler = appMockPorts.NewMockQueryHandler[query.GetPostByIdQuery, query.GetPostByIdQueryResponse](
		suite.T(),
	)

	suite.mockListQueryHandler = appMockPorts.NewMockQueryHandler[helper.Pagination, gettingpostsQuery.GetPostsQueryResponse](
		suite.T(),
	)

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
			suite.mockCommandHandler,
			suite.mockQueryHandler,
			suite.mockListQueryHandler,
			suite.mockLogger,
			suite.mockAdapter,
		)

		// Then
		suite.NotNil(postModule)
		suite.NotNil(postModule.Repository())
		suite.NotNil(postModule.CreatePostCommandHandler())
		suite.NotNil(postModule.GetPostByIdQueryHandler())
		suite.NotNil(postModule.GetPostsQueryHandler())
		suite.NotNil(postModule.Notifier())
		// Logger can be nil for testing purposes
	})
}

func (suite *PostModuleTestSuite) TestPostModule_Repository() {
	suite.Run("should return configured repository", func() {
		// Given
		postModule := NewPostModule(
			suite.mockRepo,
			suite.mockCommandHandler,
			suite.mockQueryHandler,
			suite.mockListQueryHandler,
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

func (suite *PostModuleTestSuite) TestPostModule_CommandHandler() {
	suite.Run("should return configured query handler", func() {
		// Given
		postModule := NewPostModule(
			suite.mockRepo,
			suite.mockCommandHandler,
			suite.mockQueryHandler,
			suite.mockListQueryHandler,
			suite.mockLogger,
			suite.mockAdapter,
		)

		// When
		commandHandler := postModule.CreatePostCommandHandler()

		// Then
		suite.NotNil(commandHandler)
		suite.Equal(suite.mockCommandHandler, commandHandler)
	})
}

func (suite *PostModuleTestSuite) TestPostModule_QueryHandler() {
	suite.Run("should return configured query handler", func() {
		// Given
		postModule := NewPostModule(
			suite.mockRepo,
			suite.mockCommandHandler,
			suite.mockQueryHandler,
			suite.mockListQueryHandler,
			suite.mockLogger,
			suite.mockAdapter,
		)

		// When
		queryHandler := postModule.GetPostByIdQueryHandler()

		// Then
		suite.NotNil(queryHandler)
		suite.Equal(suite.mockQueryHandler, queryHandler)
	})
}

func (suite *PostModuleTestSuite) TestPostModule_Logger() {
	suite.Run("should return configured logger", func() {
		// Given
		postModule := NewPostModule(
			suite.mockRepo,
			suite.mockCommandHandler,
			suite.mockQueryHandler,
			suite.mockListQueryHandler,
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
			suite.mockCommandHandler,
			suite.mockQueryHandler,
			suite.mockListQueryHandler,
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
		mockCommandHandler2 := appMockPorts.NewMockCommandHandler[commands.CreatePostCommandV1](
			suite.T(),
		)
		mockQueryHandler2 := appMockPorts.NewMockQueryHandler[query.GetPostByIdQuery, query.GetPostByIdQueryResponse](
			suite.T(),
		)
		mockAdapter2 := domainMockPorts.NewMockNotificationAdapter(suite.T())

		// When
		module1 := NewPostModule(
			suite.mockRepo,
			suite.mockCommandHandler,
			suite.mockQueryHandler,
			suite.mockListQueryHandler,
			suite.mockLogger,
			suite.mockAdapter,
		)
		module2 := NewPostModule(
			mockRepo2,
			mockCommandHandler2,
			mockQueryHandler2,
			suite.mockListQueryHandler,
			suite.mockLogger,
			mockAdapter2,
		)

		// Then
		suite.NotNil(module1)
		suite.NotNil(module2)
		suite.NotSame(module1, module2)
		suite.NotSame(module1.Repository(), module2.Repository())
		suite.NotSame(module1.CreatePostCommandHandler(), module2.CreatePostCommandHandler())
		suite.NotSame(module1.GetPostByIdQueryHandler(), module2.GetPostByIdQueryHandler())
		suite.NotSame(module1.Notifier(), module2.Notifier())
	})
}

func (suite *PostModuleTestSuite) TestPostModule_DependencyInjection() {
	suite.Run("should properly inject all dependencies", func() {
		// When
		postModule := NewPostModule(
			suite.mockRepo,
			suite.mockCommandHandler,
			suite.mockQueryHandler,
			suite.mockListQueryHandler,
			suite.mockLogger,
			suite.mockAdapter,
		)

		// Then
		suite.Equal(suite.mockRepo, postModule.Repository())
		suite.Equal(suite.mockCommandHandler, postModule.CreatePostCommandHandler())
		suite.Equal(suite.mockQueryHandler, postModule.GetPostByIdQueryHandler())
		suite.Equal(suite.mockLogger, postModule.Logger())
		suite.Equal(suite.mockAdapter, postModule.Notifier())
	})
}
