package domain

import "context"

type NotificationRepo interface {
	Save(ctx context.Context, i *Notification) error
}
