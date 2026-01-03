package module

import (
	postService "github.com/racibaz/go-arch/internal/modules/post/application/ports"
	postDomainPorts "github.com/racibaz/go-arch/internal/modules/post/domain/ports"
	"github.com/racibaz/go-arch/pkg/logger"
)

type PostModule struct {
	repository postDomainPorts.PostRepository
	service    postService.PostService
	logger     logger.Logger
	notifier   postDomainPorts.NotificationAdapter
}

func NewPostModule(
	repository postDomainPorts.PostRepository,
	service postService.PostService,
	logger logger.Logger,
	notifier postDomainPorts.NotificationAdapter,
) *PostModule {

	return &PostModule{
		repository: repository,
		service:    service,
		logger:     logger,
		notifier:   notifier,
	}
}

func (m PostModule) Repository() postDomainPorts.PostRepository {
	return m.repository
}

func (m PostModule) Service() postService.PostService {
	return m.service
}

func (m PostModule) Notifier() postDomainPorts.NotificationAdapter {
	return m.notifier
}

func (m PostModule) Logger() logger.Logger {
	return m.logger
}
