package module

import (
	"github.com/gin-gonic/gin"
	"github.com/racibaz/go-arch/internal/modules/post/application/commands"
	"github.com/racibaz/go-arch/internal/modules/post/application/handlers"
	postService "github.com/racibaz/go-arch/internal/modules/post/application/ports"
	postPorts "github.com/racibaz/go-arch/internal/modules/post/domain/ports"
	"github.com/racibaz/go-arch/internal/modules/post/infrastructure/messaging/rabbitmq"
	"github.com/racibaz/go-arch/internal/modules/post/infrastructure/notification/sms"
	gromPostRepo "github.com/racibaz/go-arch/internal/modules/post/infrastructure/persistence/gorm/repositories"
	"github.com/racibaz/go-arch/internal/modules/post/logging"
	"github.com/racibaz/go-arch/pkg/ddd"
	"github.com/racibaz/go-arch/pkg/logger"
	rabbitmqConn "github.com/racibaz/go-arch/pkg/messaging/rabbitmq"
)

type PostModule struct {
	repository postPorts.PostRepository
	service    postService.PostService
	logger     logger.Logger
	notifier   postPorts.NotificationAdapter
}

func NewPostModule() *PostModule {

	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	ctx := *gin.Context

	//repo := in_memory.New()
	repo := gromPostRepo.New()         // Use GORM repository for persistence
	logger, _ := logger.NewZapLogger() // Assuming NewZapLogger is a function that initializes a logger
	rabbitmqConn := rabbitmqConn.Connect()
	//messagePublisher := rabbitmq.NewMessagePublisher(config.RabbitMQ.Url, logger)
	messagePublisher := rabbitmq.NewPostMessagePublisher(rabbitmqConn, logger)
	/* todo we need to use processor in services to publish events after transaction is committed
	for now we will use directly the publisher in the service
	*/
	createPostService := commands.NewCreatePostService(repo, logger, messagePublisher)

	notificationAdapter := sms.NewTwilioSmsNotificationAdapter()

	notificationHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
		handlers.NewNotificationHandlers(notificationAdapter),
		"Notification", logger,
	)

	handlers.RegisterNotificationHandlers(notificationHandlers, domainDispatcher)

	return &PostModule{
		repository: repo,
		service:    createPostService,
		logger:     logger,
		notifier:   notificationAdapter,
	}
}

func (m PostModule) Repository() postPorts.PostRepository {
	return m.repository
}

func (m PostModule) Service() postService.PostService {
	return m.service
}

func (m PostModule) Notifier() postPorts.NotificationAdapter {
	return m.notifier
}

func (m PostModule) Logger() logger.Logger {
	return m.logger
}
