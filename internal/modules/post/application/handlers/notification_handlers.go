package handlers

import (
	"context"

	. "github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/internal/modules/post/domain/ports"
	"github.com/racibaz/go-arch/pkg/ddd"
)

type NotificationHandlers[T ddd.AggregateEvent] struct {
	notifications ports.NotificationAdapter
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*NotificationHandlers[ddd.AggregateEvent])(nil)

func NewNotificationHandlers(
	notifications ports.NotificationAdapter,
) *NotificationHandlers[ddd.AggregateEvent] {
	return &NotificationHandlers[ddd.AggregateEvent]{
		notifications: notifications,
	}
}

func (h NotificationHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case PostCreatedEvent:
		return h.onPostCreated(ctx, event)
	}
	return nil
}

func (h NotificationHandlers[T]) onPostCreated(
	ctx context.Context,
	event ddd.AggregateEvent,
) error {
	postCreated := event.Payload().(*PostCreated)
	return h.notifications.NotifyPostCreated(ctx, event.AggregateID(), postCreated.Post.UserID)
}
