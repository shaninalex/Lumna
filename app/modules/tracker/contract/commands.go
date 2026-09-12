package contract

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/bus"
)

type CreateScope struct {
	Title     string
	ProjectId int
}

func (CreateScope) Permission() (action string, scope int) { return "", 0 }

func ExecCreateScope(ctx context.Context, a *core.App, cmd CreateScope) (CreateScopeView, error) {
	return bus.Execute[CreateScope, CreateScopeView](ctx, a.Commands, cmd)
}
