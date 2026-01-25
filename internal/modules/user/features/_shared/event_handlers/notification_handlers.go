package event_handlers

import (
	"context"

	. "github.com/racibaz/go-arch/internal/modules/user/domain"
	"github.com/racibaz/go-arch/internal/modules/user/domain/ports"
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
	case UserRegisteredEvent:
		return h.onUserCreated(ctx, event)
	}
	// todo add more cases when needed
	return nil
}

func (h NotificationHandlers[T]) onUserCreated(
	ctx context.Context,
	event ddd.AggregateEvent,
) error {
	userRegistered := event.Payload().(*UserRegistered)
	return h.notifications.NotifyUserRegistered(ctx, userRegistered.User.ID())
}
