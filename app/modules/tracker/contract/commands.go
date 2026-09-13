package contract

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/bus"
)

type CreateScope struct {
	Name        string
	Description string
	ProjectId   int
}

func (CreateScope) Permission() (action string, scope int) { return "", 0 }

func ExecCreateScope(ctx context.Context, a *core.App, cmd CreateScope) (ScopeView, error) {
	return bus.Execute[CreateScope, ScopeView](ctx, a.Commands, cmd)
}

type CreateStage struct {
	ScopeID     int
	Name        string
	Description string
	Category    string
	Position    float64
	WIPLimit    *int
}

func (CreateStage) Permission() (action string, scope int) { return "", 0 }

func ExecCreateStage(ctx context.Context, a *core.App, cmd CreateStage) (StageView, error) {
	return bus.Execute[CreateStage, StageView](ctx, a.Commands, cmd)
}
