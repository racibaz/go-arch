package messaging

import (
	"context"

	postDomain "github.com/racibaz/go-arch/internal/modules/post/domain"
	userDomain "github.com/racibaz/go-arch/internal/modules/user/domain"
)

const (
	DefaultExchange    = "default_exchange"
	DeadLetterExchange = "dlx"
)

const (
	PostProcessingQueue = "post_process"
)

const (
	PostEventCreated   = "post.event.created"
	PostEventUpdated   = "post.event.updated"
	PostEventDeleted   = "post.event.deleted"
	PostEventPublished = "post.event.published"

	UserEventCreated             = "user.event.registered"
	UserEventUpdated             = "user.event.updated"
	UserEventDeleted             = "user.event.deleted"
	UserEventPublished           = "user.event.published"
	UserEventLoggedIn            = "user.event.loggedIn"
	UserEventLoggedOut           = "user.event.loggedOut"
	UserEventRefreshTokenUpdated = "user.event.refreshTokenUpdated"
)

type MessagePayload struct {
	OwnerID string `json:"ownerId"`
	Data    []byte `json:"data"`
}

type PostMessagePublisher interface {
	PublishPostCreated(ctx context.Context, payload *postDomain.Post) error
}

type UserMessagePublisher interface {
	PublishUserRegistered(ctx context.Context, payload *userDomain.User) error
	PublishUserLoggedIn(ctx context.Context, payload *userDomain.User) error
	PublishUserLoggedOut(ctx context.Context, payload *userDomain.User) error
	PublishUserRefreshTokenUpdated(ctx context.Context, payload *userDomain.User) error
}

type MessageProcessor interface {
	ProcessMessage(message any) error
}
