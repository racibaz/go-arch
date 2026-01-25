package module

import (
	"sync"

	"github.com/gin-gonic/gin"
	eventHandler "github.com/racibaz/go-arch/internal/modules/user/features/_shared/event_handlers"
	registeringUserV1Endpoint "github.com/racibaz/go-arch/internal/modules/user/features/registering/v1/adapters/endpoints"
	commandsV1Endpoint "github.com/racibaz/go-arch/internal/modules/user/features/registering/v1/application/commands"
	"github.com/racibaz/go-arch/internal/modules/user/infrastructure/messaging/rabbitmq"
	"github.com/racibaz/go-arch/internal/modules/user/infrastructure/notification/sms"
	"github.com/racibaz/go-arch/internal/modules/user/infrastructure/observability/logging"
	gormUserRepo "github.com/racibaz/go-arch/internal/modules/user/infrastructure/persistence/gorm/repositories"
	"github.com/racibaz/go-arch/pkg/ddd"
	"github.com/racibaz/go-arch/pkg/logger"
	rabbitmqConn "github.com/racibaz/go-arch/pkg/messaging/rabbitmq"
	googleGrpc "google.golang.org/grpc"
)

var (
	userModuleInstance *UserModule
	once               sync.Once
)

func BuildModule() *UserModule {
	// Return existing instance if already created
	if userModuleInstance != nil {
		return userModuleInstance
	}

	// Create the instance only once
	once.Do(func() {
		repository := gormUserRepo.NewGormUserRepository()

		// Assuming NewZapLogger is a function that initializes a logger
		logger, _ := logger.NewZapLogger()

		domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()

		rabbitmqConn := rabbitmqConn.Connection()

		messagePublisher := rabbitmq.NewUserMessagePublisher(rabbitmqConn, logger)
		/* todo we need to use processor in handler to publish events after transaction is committed
		for now we will use directly the publisher in the handler
		*/
		registerUserCommandHandler := commandsV1Endpoint.NewRegisterUserHandler(
			repository,
			logger,
			messagePublisher,
		)

		notificationAdapter := sms.NewTwilioSmsNotificationAdapter()

		notificationHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
			eventHandler.NewNotificationHandlers(notificationAdapter),
			"Notification", logger,
		)

		RegisterNotificationHandlers(notificationHandlers, domainDispatcher)

		userModuleInstance = NewUserModule(
			repository,
			registerUserCommandHandler,
			logger,
			notificationAdapter,
		)
	})
	return userModuleInstance
}

func Routes(router *gin.Engine) {
	module := BuildModule()

	// Collect here restful routes of your module.
	registeringUserV1Endpoint.MapHttpRoute(router, module.RegisterUserCommandHandler())
}

func GrpcRoutes(grpcServer *googleGrpc.Server) {
	// Collect here grpc routes of your module
}
