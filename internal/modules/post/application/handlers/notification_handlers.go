package handlers

import (
	"context"
	"github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/internal/modules/post/domain/ports"
	"github.com/racibaz/go-arch/pkg/ddd"
)

type NotificationHandlers struct {
	notifications ports.NotificationRepository
	ignoreUnimplementedDomainEvents
}

var _ DomainEventHandlers = (*NotificationHandlers)(nil)

func NewNotificationHandlers(notifications ports.NotificationRepository) *NotificationHandlers {
	return &NotificationHandlers{
		notifications: notifications,
	}
}

func (h NotificationHandlers) OnPostCreated(ctx context.Context, event ddd.Event) error {
	postCreated := event.(*domain.PostCreated)
	return h.notifications.NotifyPostCreated(ctx, postCreated.Post.ID, postCreated.Post.UserID)
}
