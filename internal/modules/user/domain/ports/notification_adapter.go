package ports

import (
	"context"
)

type NotificationAdapter interface {
	NotifyUserRegistered(ctx context.Context, UserID string) error
	NotifyUserCanceled(ctx context.Context, UserID string) error
	NotifyUserReady(ctx context.Context, UserID string) error
}
