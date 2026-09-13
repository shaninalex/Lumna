package contract

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/bus"
)

type WorkspaceList struct{}

func (WorkspaceList) Permission() (action string, scope int) { return "", 0 }

func AskWorkspaceList(ctx context.Context, a *core.App, q WorkspaceList) ([]WorkspaceView, error) {
	return bus.Ask[WorkspaceList, []WorkspaceView](ctx, a.Queries, q)
}

// CreateWorkspace — Creates workspace
type CreateWorkspace struct {
	Title      string
	OwnerEmail string
	Active     bool
}

func (CreateWorkspace) Permission() (action string, scope int) { return "", 0 }

func ExecCreateWorkspace(ctx context.Context, a *core.App, c CreateWorkspace) (WorkspaceView, error) {
	return bus.Execute[CreateWorkspace, WorkspaceView](ctx, a.Commands, c)
}

type AddIdentityToWorkspace struct {
	IdentityId  int
	WorkspaceId int
}

func (AddIdentityToWorkspace) Permission() (action string, scope int) { return "", 0 }

func ExecAddIdentityToWorkspace(ctx context.Context, a *core.App, cmd AddIdentityToWorkspace) (bool, error) {
	return bus.Execute[AddIdentityToWorkspace, bool](ctx, a.Commands, cmd)
}
