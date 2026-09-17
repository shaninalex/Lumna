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

type StageCreate struct {
	ScopeID     int
	Name        string
	Description string
	Category    string
	Position    float64
	WIPLimit    *int
}

func (StageCreate) Permission() (action string, scope int) { return "", 0 }

func ExecCreateStage(ctx context.Context, a *core.App, cmd StageCreate) (StageView, error) {
	return bus.Execute[StageCreate, StageView](ctx, a.Commands, cmd)
}

type WorkItemCreate struct {
	Title       string
	Description *string
	ProjectId   int
	Position    float64
	StageId     *int
	ScopeId     *int
}

func (WorkItemCreate) Permission() (action string, scope int) { return "", 0 }

func ExecWorkItemCreate(ctx context.Context, a *core.App, cmd WorkItemCreate) (WorkItemView, error) {
	return bus.Execute[WorkItemCreate, WorkItemView](ctx, a.Commands, cmd)
}

// WorkItemMove - a post request data described moving item in a single column
// TODO: rename properties and tags
type WorkItemMove struct {
	Rank       float64 `json:"rank"`
	ScopeId    int     `json:"board_id"`
	WorkItemId int     `json:"task_id"`
}

func (WorkItemMove) Permission() (action string, scope int) { return "", 0 }

func ExecWorkItemMove(ctx context.Context, a *core.App, cmd WorkItemMove) (WorkItemView, error) {
	return bus.Execute[WorkItemMove, WorkItemView](ctx, a.Commands, cmd)
}
