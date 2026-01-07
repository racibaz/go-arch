package module

import (
	"github.com/racibaz/go-arch/internal/modules/post/application/command"
	ports "github.com/racibaz/go-arch/internal/modules/post/application/ports"
	"github.com/racibaz/go-arch/internal/modules/post/application/query"
	postDomainPorts "github.com/racibaz/go-arch/internal/modules/post/domain/ports"
	"github.com/racibaz/go-arch/pkg/logger"
)

// PostModule encapsulates the components related to the Post module.
type PostModule struct {
	repository               postDomainPorts.PostRepository
	createPostCommandHandler ports.CommandHandler[command.CreatePostCommand]
	getPostQueryHandler      ports.QueryHandler[query.GetPostByIdQuery, query.GetPostByIdQueryResponse]
	logger                   logger.Logger
	notifier                 postDomainPorts.NotificationAdapter
}

// NewPostModule initializes a new PostModule with the provided components.
func NewPostModule(
	repository postDomainPorts.PostRepository,
	createPostCommandHandler ports.CommandHandler[command.CreatePostCommand],
	getPostQueryHandler ports.QueryHandler[query.GetPostByIdQuery, query.GetPostByIdQueryResponse],
	logger logger.Logger,
	notifier postDomainPorts.NotificationAdapter,
) *PostModule {
	return &PostModule{
		repository:               repository,
		createPostCommandHandler: createPostCommandHandler,
		getPostQueryHandler:      getPostQueryHandler,
		logger:                   logger,
		notifier:                 notifier,
	}
}

func (m PostModule) Repository() postDomainPorts.PostRepository {
	return m.repository
}

func (m PostModule) CommandHandler() ports.CommandHandler[command.CreatePostCommand] {
	return m.createPostCommandHandler
}

func (m PostModule) QueryHandler() ports.QueryHandler[query.GetPostByIdQuery, query.GetPostByIdQueryResponse] {
	return m.getPostQueryHandler
}

func (m PostModule) Notifier() postDomainPorts.NotificationAdapter {
	return m.notifier
}

func (m PostModule) Logger() logger.Logger {
	return m.logger
}
