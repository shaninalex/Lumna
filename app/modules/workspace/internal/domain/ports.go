package domain

import (
	"context"
)

type WorkspaceRepo interface {
	Save(ctx context.Context, w Workspace) (Workspace, error)
	List(ctx context.Context) ([]Workspace, error)
	Delete(ctx context.Context, id int) error
}

type IdentityWorkspaceRepo interface {
	Save(ctx context.Context, i IdentityWorkspace) (IdentityWorkspace, error)
	ByWorkspaceId(ctx context.Context, workspaceId int) ([]IdentityWorkspace, error)
	Delete(ctx context.Context, identityID, workspaceID int) error
}

type ProjectRepo interface {
	Save(ctx context.Context, p Project) (Project, error)
	ByWorkspaceId(ctx context.Context, workspaceId int) ([]Project, error)
	Update(ctx context.Context, t Project) (Project, error)
	GetById(ctx context.Context, projectID int) (Project, error)
}

type InvitationRepo interface {
	Create(ctx context.Context, i Invitation) (Invitation, error)
	GetByHash(ctx context.Context, hash string) (Invitation, error)
	GetById(ctx context.Context, invitationId int) (Invitation, error)
	ListByWorkspaceId(ctx context.Context, workspaceId int) ([]Invitation, error)
	Delete(ctx context.Context, invitationId int) error
	Update(ctx context.Context, i Invitation) error
}

type TokenHasher interface {
	CreateToken() (string, string, error)
	Hash(token string) string
}
