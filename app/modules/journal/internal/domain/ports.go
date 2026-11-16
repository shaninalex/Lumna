package domain

import "context"

type EntryRepo interface {
	Save(ctx context.Context, e Entry) (Entry, error)
}
