package ports

import (
	"context"

	"github.com/racibaz/go-arch/pkg/ddd"
)

// DomainEventHandlers defines the interface for handling domain events related to posts.
type DomainEventHandlers interface {
	OnPostCreated(ctx context.Context, event ddd.Event) error
	OnPostReadied(ctx context.Context, event ddd.Event) error
	OnPostCanceled(ctx context.Context, event ddd.Event) error
	OnPostCompleted(ctx context.Context, event ddd.Event) error
}

type ignoreUnimplementedDomainEvents struct{}

var _ DomainEventHandlers = (*ignoreUnimplementedDomainEvents)(nil)

func (ignoreUnimplementedDomainEvents) OnPostCreated(ctx context.Context, event ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnPostReadied(ctx context.Context, event ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnPostCanceled(ctx context.Context, event ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnPostCompleted(ctx context.Context, event ddd.Event) error {
	return nil
}
