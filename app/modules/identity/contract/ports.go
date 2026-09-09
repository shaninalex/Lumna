package contract

import "context"

type Reader interface {
	DisplayNames(ctx context.Context, ids []int) (map[int]string, error)
	Exists(ctx context.Context, id int) (bool, error)
}

type Mailer interface {
	Send(ctx context.Context, to, template string, data map[string]any) error
}
