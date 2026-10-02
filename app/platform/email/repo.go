package email

import "context"

type Repository interface {
	PendingEmails() ([]string, error)
	Save(ctx context.Context, emails []string) error
}

type Writer interface {
	Send(ctx context.Context, e Email) error
}
