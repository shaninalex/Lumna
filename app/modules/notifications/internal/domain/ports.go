package domain

import "context"

type NotificationRepo interface {
	Save(ctx context.Context, n Notification) (Notification, error)
}

type EmailRepo interface {
	PendingEmails(ctx context.Context) ([]Email, error)
	Create(ctx context.Context, entry Email) error
	Update(ctx context.Context, entry Email) error
}

type EntryContentBuilder interface {
	GenerateTemplates(ctx context.Context) ([]byte, []byte)
}
