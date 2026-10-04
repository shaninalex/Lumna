package domain

import (
	"context"
)

type WorkspaceRepo interface {
	Save(ctx context.Context, w Workspace) (Workspace, error)
	List(ctx context.Context) ([]Workspace, error)
}

type IdentityWorkspaceRepo interface {
	Save(ctx context.Context, i IdentityWorkspace) (IdentityWorkspace, error)
	ByWorkspaceId(ctx context.Context, workspaceId int) ([]IdentityWorkspace, error)
	//ByIdentityId(ctx context.Context, identityId int) ([]IdentityWorkspace, error)
	//Delete(ctx context.Context, identityId, workspaceId int) (bool, error)
}

type ProjectRepo interface {
	Save(ctx context.Context, p Project) (Project, error)
	ByWorkspaceId(ctx context.Context, workspaceId int) ([]Project, error)
}

type InvitationRepo interface {
	Create(ctx context.Context, i Invitation) (Invitation, error)
	GetByHash(ctx context.Context, hash string) (Invitation, error)
	ListByWorkspaceId(ctx context.Context, workspaceId int) ([]Invitation, error)
}

type TokenHasher interface {
	CreateToken() (string, string, error)
}
