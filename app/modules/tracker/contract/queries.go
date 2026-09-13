package contract

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/bus"
)

type ScopeList struct {
	ProjectId int
}

func (ScopeList) Permission() (action string, scope int) { return "", 0 }

func AskScopeList(ctx context.Context, a *core.App, q ScopeList) ([]ScopeView, error) {
	return bus.Ask[ScopeList, []ScopeView](ctx, a.Queries, q)
}

type StageList struct {
	ScopeId int
}

func (StageList) Permission() (action string, scope int) { return "", 0 }

func AskStageList(ctx context.Context, a *core.App, q StageList) ([]StageView, error) {
	return bus.Ask[StageList, []StageView](ctx, a.Queries, q)
}
