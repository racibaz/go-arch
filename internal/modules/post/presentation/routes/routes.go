package routes

import (
	"sync"

	"github.com/gin-gonic/gin"
	postModule "github.com/racibaz/go-arch/internal/modules/post"
	"github.com/racibaz/go-arch/internal/modules/post/application/commands"
	"github.com/racibaz/go-arch/internal/modules/post/application/handlers"
	"github.com/racibaz/go-arch/internal/modules/post/infrastructure/messaging/rabbitmq"
	"github.com/racibaz/go-arch/internal/modules/post/infrastructure/notification/sms"
	gormPostRepo "github.com/racibaz/go-arch/internal/modules/post/infrastructure/persistence/gorm/repositories"
	"github.com/racibaz/go-arch/internal/modules/post/logging"
	postGrpcController "github.com/racibaz/go-arch/internal/modules/post/presentation/grpc"
	postController "github.com/racibaz/go-arch/internal/modules/post/presentation/http"
	"github.com/racibaz/go-arch/pkg/ddd"
	"github.com/racibaz/go-arch/pkg/logger"
	rabbitmqConn "github.com/racibaz/go-arch/pkg/messaging/rabbitmq"
	googleGrpc "google.golang.org/grpc"
)

var (
	postModuleInstance *postModule.PostModule
	once               sync.Once
)

func BuildPostModule() *postModule.PostModule {
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
		/* todo we need to use processor in services to publish events after transaction is committed
		for now we will use directly the publisher in the service
		*/
		createPostService := commands.NewCreatePostService(repository, logger, messagePublisher)

		notificationAdapter := sms.NewTwilioSmsNotificationAdapter()

		notificationHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
			handlers.NewNotificationHandlers(notificationAdapter),
			"Notification", logger,
		)

		handlers.RegisterNotificationHandlers(notificationHandlers, domainDispatcher)

		postModuleInstance = postModule.NewPostModule(
			repository,
			createPostService,
			logger,
			notificationAdapter,
		)
	})
	return postModuleInstance
}

func Routes(router *gin.Engine) {
	module := BuildPostModule()
	newPostController := postController.NewPostController(module.Service())

	v1 := router.Group("/api/v1")
	{
		eg := v1.Group("/posts")
		{
			eg.GET("/:id", newPostController.Show)
			eg.POST("/", newPostController.Store)
		}
	}
}

func GrpcRoutes(grpcServer *googleGrpc.Server) {
	module := BuildPostModule()

	postGrpcController.NewPostGrpcController(grpcServer, module.Service())
}
