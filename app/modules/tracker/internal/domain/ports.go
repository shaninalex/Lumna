package domain

import "context"

type ScopeRepo interface {
	Create(ctx context.Context, scope Scope) (Scope, error)
	Update(ctx context.Context, scope Scope) error
	ListByProject(ctx context.Context, projectId int) ([]Scope, error)
}

type StageRepo interface {
	Create(ctx context.Context, stage Stage) (Stage, error)
	Update(ctx context.Context, stage Stage) error
	Get(ctx context.Context, stageId int) (Stage, error)
	ListByScope(ctx context.Context, scopeId int) ([]Stage, error)
	Delete(ctx context.Context, stageId int) error
}

type WorkingItemRepo interface {
	Create(ctx context.Context, wi WorkItem) (WorkItem, error)
	Update(ctx context.Context, wi WorkItem) error
	Get(ctx context.Context, itemId int) (WorkItem, error)
	ListByScope(ctx context.Context, scopeId int) ([]WorkItem, error)
	Assignment(ctx context.Context, identity, itemId int) error
	Delete(ctx context.Context, itemId int) error
	BatchDelete(ctx context.Context, itemIds []int) error
}
