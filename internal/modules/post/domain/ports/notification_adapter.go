package ports

import (
	"context"
)

type NotificationAdapter interface {
	NotifyPostCreated(ctx context.Context, postID, UserID string) error
	NotifyPostCanceled(ctx context.Context, postID, UserID string) error
	NotifyPostReady(ctx context.Context, postID, UserID string) error
}
