package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
)

type CreateWorkspace struct {
	workspaceRepo domain.WorkspaceRepo
}

func NewCreateWorkspace(workspaceRepo domain.WorkspaceRepo) *CreateWorkspace {
	return &CreateWorkspace{
		workspaceRepo: workspaceRepo,
	}
}

func (s *CreateWorkspace) Handle(ctx context.Context, cmd contract.CreateWorkspace) (contract.WorkspaceView, error) {
	workspace := domain.NewWorkspace(cmd.Title, cmd.OwnerEmail, cmd.Active)
	if err := s.workspaceRepo.Save(ctx, workspace); err != nil {
		return contract.WorkspaceView{}, err
	}
	return contract.WorkspaceView{
		Id:         workspace.ID,
		Title:      workspace.Title,
		OwnerEmail: workspace.OwnerEmail,
		Active:     workspace.Active,
		CreatedAt:  workspace.CreatedAt,
		UpdatedAt:  workspace.UpdatedAt,
	}, nil
}
