package domain

import "context"

type ScopeRepo interface {
	Save(ctx context.Context, scope *Scope) error
	Get(ctx context.Context, projectId int) ([]Scope, error)
}
