package logging

import (
	"context"
	"github.com/racibaz/go-arch/internal/modules/post/application/handlers"
	"github.com/racibaz/go-arch/pkg/ddd"
	"github.com/racibaz/go-arch/pkg/logger"
)

type DomainEventHandlers struct {
	handlers.DomainEventHandlers
	logger logger.Logger
}

var _ handlers.DomainEventHandlers = (*DomainEventHandlers)(nil)

func LogDomainEventHandlerAccess(handlers handlers.DomainEventHandlers, logger logger.Logger) DomainEventHandlers {
	return DomainEventHandlers{
		DomainEventHandlers: handlers,
		logger:              logger,
	}
}

func (h DomainEventHandlers) OnPostCreated(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info("--> Post.OnPostCreated")
	defer func() { h.logger.Error("<-- Post.OnPostCreated") }()
	return h.DomainEventHandlers.OnPostCreated(ctx, event)
}
