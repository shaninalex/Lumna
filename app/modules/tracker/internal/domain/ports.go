package domain

import "context"

type ScopeRepo interface {
	Save(ctx context.Context, scope *Scope) error
	Get(ctx context.Context, projectId int) ([]Scope, error)
}

type StageRepo interface {
	Save(ctx context.Context, scope *Stage) error
	Get(ctx context.Context, scopeId int) ([]Stage, error)
	GetById(ctx context.Context, stageId int) (*Stage, error)
}

type WorkingItemRepo interface {
	Save(ctx context.Context, wi *WorkItem) error
	List(ctx context.Context, scopeId int) ([]WorkItem, error)
	Get(ctx context.Context, itemId int) (*WorkItem, error)
	Assignment(ctx context.Context, identity, itemId int) error
	Delete(ctx context.Context, itemId int) (bool, error)
}
