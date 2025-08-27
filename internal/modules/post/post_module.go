package post_module

import (
	"github.com/racibaz/go-arch/internal/modules/post/application/handlers"
	postService "github.com/racibaz/go-arch/internal/modules/post/application/ports"
	"github.com/racibaz/go-arch/internal/modules/post/application/usecases"
	postPorts "github.com/racibaz/go-arch/internal/modules/post/domain/ports"
	"github.com/racibaz/go-arch/internal/modules/post/infrastructure/notification/sms"
	gromPostRepo "github.com/racibaz/go-arch/internal/modules/post/infrastructure/persistence/gorm/repositories"
	"github.com/racibaz/go-arch/internal/modules/post/logging"
	"github.com/racibaz/go-arch/pkg/ddd"
	"github.com/racibaz/go-arch/pkg/logger"
)

type PostModule struct {
	repository postPorts.PostRepository
	service    postService.PostService
	logger     logger.Logger
	notifier   postPorts.NotificationAdapter
}

func NewPostModule() *PostModule {

	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()

	//repo := in_memory.New()
	repo := gromPostRepo.New()         // Use GORM repository for persistence
	logger, _ := logger.NewZapLogger() // Assuming NewZapLogger is a function that initializes a logger
	service := usecases.NewPostUseCase(repo, logger)
	notificationAdapter := sms.NewTwilioSmsNotificationAdapter()

	notificationHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
		handlers.NewNotificationHandlers(notificationAdapter),
		"Notification", logger,
	)

	handlers.RegisterNotificationHandlers(notificationHandlers, domainDispatcher)

	return &PostModule{
		repository: repo,
		service:    service,
		logger:     logger,
		notifier:   notificationAdapter,
	}
}

func (m PostModule) GetRepository() postPorts.PostRepository {
	return m.repository
}

func (m PostModule) GetService() postService.PostService {
	return m.service
}

func (m PostModule) GetNotifier() postPorts.NotificationAdapter {
	return m.notifier
}

func (m PostModule) GetLogger() logger.Logger {
	return m.logger
}
