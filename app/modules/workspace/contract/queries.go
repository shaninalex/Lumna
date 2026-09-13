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

type ProjectList struct {
	WorkspaceId int
}

func (ProjectList) Permission() (action string, scope int) { return "", 0 }

func AskProjectList(ctx context.Context, a *core.App, q ProjectList) ([]ProjectView, error) {
	return bus.Ask[ProjectList, []ProjectView](ctx, a.Queries, q)
}
