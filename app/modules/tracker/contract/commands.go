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
