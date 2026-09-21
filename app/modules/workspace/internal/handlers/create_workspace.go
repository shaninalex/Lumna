package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type CreateWorkspace struct {
	workspaceRepo domain.WorkspaceRepo
	clock         clock.Clock
}

func NewCreateWorkspace(workspaceRepo domain.WorkspaceRepo, clk clock.Clock) *CreateWorkspace {
	return &CreateWorkspace{
		workspaceRepo: workspaceRepo,
		clock:         clk,
	}
}

func (s *CreateWorkspace) Handle(ctx context.Context, cmd contract.CreateWorkspace) (contract.WorkspaceView, error) {
	workspace, err := s.workspaceRepo.Save(ctx, domain.NewWorkspace(cmd.Title, cmd.OwnerEmail, cmd.Active, s.clock.Now()))
	if err != nil {
		return contract.WorkspaceView{}, err
	}
	return toWorkspaceView(workspace), nil
}
