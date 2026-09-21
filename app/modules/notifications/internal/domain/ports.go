package domain

import "context"

type NotificationRepo interface {
	Save(ctx context.Context, n Notification) (Notification, error)
}
