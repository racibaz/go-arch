package messaging

import (
	"context"
	"github.com/racibaz/go-arch/internal/modules/post/domain"
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
)

type MessagePayload struct {
	OwnerID string `json:"ownerId"`
	Data    []byte `json:"data"`
}

type MessagePublisher interface {
	PublishPostCreated(ctx context.Context, payload *domain.Post) error
}

type MessageProcessor interface {
	ProcessMessage(message any) error
}
