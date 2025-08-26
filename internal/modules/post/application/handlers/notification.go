package handlers

import (
	"github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/pkg/ddd"
)

func RegisterNotificationHandlers(notificationHandlers DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.PostCreated{}, notificationHandlers.OnPostCreated)
}
