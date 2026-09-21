package domain

import "context"

type WorkspaceRepo interface {
	Save(ctx context.Context, w Workspace) (Workspace, error)
	List(ctx context.Context) ([]Workspace, error)
}

type IdentityWorkspaceRepo interface {
	Save(ctx context.Context, i IdentityWorkspace) (IdentityWorkspace, error)
	//ByIdentityId(ctx context.Context, identityId int) ([]IdentityWorkspace, error)
	//ByWorkspaceId(ctx context.Context, workspaceId int) ([]IdentityWorkspace, error)
	//Delete(ctx context.Context, identityId, workspaceId int) (bool, error)
}

type ProjectRepo interface {
	Save(ctx context.Context, p Project) (Project, error)
	ByWorkspaceId(ctx context.Context, workspaceId int) ([]Project, error)
}
