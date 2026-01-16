package module

import (
	"sync"

	"github.com/gin-gonic/gin"
	eventHandler "github.com/racibaz/go-arch/internal/modules/post/features/_shared/event_handlers"
	creatingPostV1Endpoint "github.com/racibaz/go-arch/internal/modules/post/features/creatingpost/v1/adapters/endpoints"
	"github.com/racibaz/go-arch/internal/modules/post/features/creatingpost/v1/adapters/endpoints/grpc"
	"github.com/racibaz/go-arch/internal/modules/post/features/creatingpost/v1/application/commands"
	getPostByPostByIdV1Endpoint "github.com/racibaz/go-arch/internal/modules/post/features/gettingpostbyid/v1/adapters/endpoints"
	"github.com/racibaz/go-arch/internal/modules/post/features/gettingpostbyid/v1/application/query"
	"github.com/racibaz/go-arch/internal/modules/post/infrastructure/messaging/rabbitmq"
	"github.com/racibaz/go-arch/internal/modules/post/infrastructure/notification/sms"
	"github.com/racibaz/go-arch/internal/modules/post/infrastructure/observability/logging"
	gormPostRepo "github.com/racibaz/go-arch/internal/modules/post/infrastructure/persistence/gorm/repositories"
	"github.com/racibaz/go-arch/pkg/ddd"
	"github.com/racibaz/go-arch/pkg/logger"
	rabbitmqConn "github.com/racibaz/go-arch/pkg/messaging/rabbitmq"
	googleGrpc "google.golang.org/grpc"
)

var (
	postModuleInstance *PostModule
	once               sync.Once
)

func BuildPostModule() *PostModule {
	// Return existing instance if already created
	if postModuleInstance != nil {
		return postModuleInstance
	}

	// Create the instance only once
	once.Do(func() {
		// Use In-memory for persistence
		// repo := in_memory.NewGormPostRepository()
		// Use GORM repository for persistence
		repository := gormPostRepo.NewGormPostRepository()

		// Assuming NewZapLogger is a function that initializes a logger
		logger, _ := logger.NewZapLogger()

		domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()

		rabbitmqConn := rabbitmqConn.Connection()

		messagePublisher := rabbitmq.NewPostMessagePublisher(rabbitmqConn, logger)
		/* todo we need to use processor in handler to publish events after transaction is committed
		for now we will use directly the publisher in the handler
		*/
		createPostCommandHandler := commands.NewCreatePostHandler(
			repository,
			logger,
			messagePublisher,
		)
		getPostQueryHandler := query.NewGetPostHandler(repository, logger)

		notificationAdapter := sms.NewTwilioSmsNotificationAdapter()

		notificationHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
			eventHandler.NewNotificationHandlers(notificationAdapter),
			"Notification", logger,
		)

		RegisterNotificationHandlers(notificationHandlers, domainDispatcher)

		postModuleInstance = NewPostModule(
			repository,
			createPostCommandHandler,
			getPostQueryHandler,
			logger,
			notificationAdapter,
		)
	})
	return postModuleInstance
}

func Routes(router *gin.Engine) {
	module := BuildPostModule()

	// Collect here restful routes of your module.
	creatingPostV1Endpoint.MapHttpRoute(router, module.CommandHandler())
	getPostByPostByIdV1Endpoint.MapHttpRoute(router, module.QueryHandler())
}

func GrpcRoutes(grpcServer *googleGrpc.Server) {
	module := BuildPostModule()

	// Collect here grpc routes of your module
	grpc.NewCreatePostHandler(grpcServer, module.CommandHandler())
}
