package post_module

import (
	"github.com/racibaz/go-arch/internal/modules/post/application/handlers"
	postService "github.com/racibaz/go-arch/internal/modules/post/application/ports"
	"github.com/racibaz/go-arch/internal/modules/post/application/usecases"
	postRepository "github.com/racibaz/go-arch/internal/modules/post/domain/ports"
	"github.com/racibaz/go-arch/internal/modules/post/infrastructure/notification/sms"
	gromPostRepo "github.com/racibaz/go-arch/internal/modules/post/infrastructure/persistence/gorm/repositories"
	"github.com/racibaz/go-arch/internal/modules/post/logging"
	"github.com/racibaz/go-arch/pkg/ddd"
	"github.com/racibaz/go-arch/pkg/logger"
)

type PostModule struct {
	repository postRepository.PostRepository
	service    postService.PostService
}

func NewPostModule() *PostModule {

	domainDispatcher := ddd.NewEventDispatcher()

	//repo := in_memory.New()
	repo := gromPostRepo.New()         // Use GORM repository for persistence
	logger, _ := logger.NewZapLogger() // Assuming NewZapLogger is a function that initializes a logger
	service := usecases.NewPostUseCase(repo, logger)
	notificationRepository := sms.NewTwilloSmsNotificationRepository()

	notificationHandlers := logging.LogDomainEventHandlerAccess(
		handlers.NewNotificationHandlers(notificationRepository),
		logger,
	)

	handlers.RegisterNotificationHandlers(notificationHandlers, domainDispatcher)

	return &PostModule{
		repository: repo,
		service:    service,
	}
}

func (m *PostModule) GetRepository() postRepository.PostRepository {
	return m.repository
}

func (m *PostModule) GetService() postService.PostService {
	return m.service
}
