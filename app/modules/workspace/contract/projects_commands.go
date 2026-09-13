package contract

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/bus"
)

type CreateProject struct {
	Title       string
	WorkspaceId int
	OwnerId     int
}

func (CreateProject) Permission() (action string, scope int) { return "", 0 }

func ExecCreateProject(ctx context.Context, a *core.App, cmd CreateProject) (ProjectView, error) {
	return bus.Execute[CreateProject, ProjectView](ctx, a.Commands, cmd)
}

type ProjectList struct {
	WorkspaceId int
}

func (ProjectList) Permission() (action string, scope int) { return "", 0 }

func AskProjectList(ctx context.Context, a *core.App, q ProjectList) ([]ProjectView, error) {
	return bus.Ask[ProjectList, []ProjectView](ctx, a.Queries, q)
}
