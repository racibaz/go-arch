package logging

import (
	"context"
	"github.com/racibaz/go-arch/pkg/ddd"
	"github.com/racibaz/go-arch/pkg/logger"
)

type EventHandlers[T ddd.Event] struct {
	ddd.EventHandler[T]
	label  string
	logger logger.Logger
}

var _ ddd.EventHandler[ddd.Event] = (*EventHandlers[ddd.Event])(nil)

func LogEventHandlerAccess[T ddd.Event](handlers ddd.EventHandler[T], label string, logger logger.Logger) EventHandlers[T] {
	return EventHandlers[T]{
		EventHandler: handlers,
		label:        label,
		logger:       logger,
	}
}

func (h EventHandlers[T]) HandleEvent(ctx context.Context, event T) (err error) {
	h.logger.Info("--> Post.%s.On(%s)", h.label, event.EventName())
	defer func() { h.logger.Error("<-- Post.%s.On(%s)", h.label, event.EventName()) }()
	return h.EventHandler.HandleEvent(ctx, event)
}
