package module

import (
	"github.com/racibaz/go-arch/internal/modules/post/domain"
	postDomainPorts "github.com/racibaz/go-arch/internal/modules/post/domain/ports"
	"github.com/racibaz/go-arch/internal/modules/post/features/creatingpost/v1/application/commands"
	"github.com/racibaz/go-arch/internal/modules/post/features/gettingpostbyid/v1/application/query"
	gettingposts "github.com/racibaz/go-arch/internal/modules/post/features/gettingposts/v1/application/query"
	ports2 "github.com/racibaz/go-arch/internal/modules/shared/application/ports"
	"github.com/racibaz/go-arch/pkg/ddd"
	"github.com/racibaz/go-arch/pkg/helper"
	"github.com/racibaz/go-arch/pkg/logger"
)

// PostModule encapsulates the components related to the Post module.
type PostModule struct {
	repository               postDomainPorts.PostRepository
	createPostCommandHandler ports2.CommandHandler[commands.CreatePostCommandV1]
	getPostQueryHandler      ports2.QueryHandler[query.GetPostByIdQuery, query.GetPostByIdQueryResponse]
	getPostsQueryHandler     ports2.QueryHandler[helper.Pagination, gettingposts.GetPostsQueryResponse]
	logger                   logger.Logger
	notifier                 postDomainPorts.NotificationAdapter
}

// NewPostModule initializes a new PostModule with the provided components.
func NewPostModule(
	repository postDomainPorts.PostRepository,
	createPostCommandHandler ports2.CommandHandler[commands.CreatePostCommandV1],
	getPostQueryHandler ports2.QueryHandler[query.GetPostByIdQuery, query.GetPostByIdQueryResponse],
	getPostsQueryHandler ports2.QueryHandler[helper.Pagination, gettingposts.GetPostsQueryResponse],
	logger logger.Logger,
	notifier postDomainPorts.NotificationAdapter,
) *PostModule {
	return &PostModule{
		repository:               repository,
		createPostCommandHandler: createPostCommandHandler,
		getPostQueryHandler:      getPostQueryHandler,
		getPostsQueryHandler:     getPostsQueryHandler,
		logger:                   logger,
		notifier:                 notifier,
	}
}

func (m PostModule) Repository() postDomainPorts.PostRepository {
	return m.repository
}

func (m PostModule) CreatePostCommandHandler() ports2.CommandHandler[commands.CreatePostCommandV1] {
	return m.createPostCommandHandler
}

func (m PostModule) GetPostByIdQueryHandler() ports2.QueryHandler[query.GetPostByIdQuery, query.GetPostByIdQueryResponse] {
	return m.getPostQueryHandler
}

func (m PostModule) GetPostsQueryHandler() ports2.QueryHandler[helper.Pagination, gettingposts.GetPostsQueryResponse] {
	return m.getPostsQueryHandler
}

func (m PostModule) Notifier() postDomainPorts.NotificationAdapter {
	return m.notifier
}

func (m PostModule) Logger() logger.Logger {
	return m.logger
}

// RegisterNotificationHandlers registers notification handlers for domain events.
func RegisterNotificationHandlers(
	notificationHandlers ddd.EventHandler[ddd.AggregateEvent],
	domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent],
) {
	domainSubscriber.Subscribe(domain.PostCreatedEvent, notificationHandlers)
}
