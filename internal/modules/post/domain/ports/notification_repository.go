package ports

import (
	"context"
)

type NotificationRepository interface {
	NotifyPostCreated(ctx context.Context, postID, UserID string) error
	NotifyPostCanceled(ctx context.Context, postID, UserID string) error
	NotifyPostReady(ctx context.Context, postID, UserID string) error
}
