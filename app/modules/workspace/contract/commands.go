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
	Role        string
}

func (AddIdentityToWorkspace) Permission() (action string, scope int) { return "", 0 }

func ExecAddIdentityToWorkspace(ctx context.Context, a *core.App, cmd AddIdentityToWorkspace) (bool, error) {
	return bus.Execute[AddIdentityToWorkspace, bool](ctx, a.Commands, cmd)
}

type CreateInvitation struct {
	Invitation InvitationView
}

func (CreateInvitation) Permission() (action string, scope int) { return "admin", 0 }

func ExecCreateInvitation(ctx context.Context, a *core.App, cmd CreateInvitation) (InvitationView, error) {
	return bus.Execute[CreateInvitation, InvitationView](ctx, a.Commands, cmd)
}

type DeleteInvitation struct {
	InvitationId int
}

func (DeleteInvitation) Permission() (action string, scope int) { return "", 0 }

func ExecDeleteInvitation(ctx context.Context, a *core.App, cmd DeleteInvitation) (bool, error) {
	return bus.Execute[DeleteInvitation, bool](ctx, a.Commands, cmd)
}

type RevokeInvitation struct {
	InvitationId int
}

func (RevokeInvitation) Permission() (action string, scope int) { return "", 0 }

func ExecRevokeInvitation(ctx context.Context, a *core.App, cmd RevokeInvitation) (InvitationView, error) {
	return bus.Execute[RevokeInvitation, InvitationView](ctx, a.Commands, cmd)
}
