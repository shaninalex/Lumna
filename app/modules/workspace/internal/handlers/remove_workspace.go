package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
)

type RemoveWorkspace struct {
	workspaceRepo domain.WorkspaceRepo
}

func NewRemoveWorkspace(wr domain.WorkspaceRepo) *RemoveWorkspace {
	return &RemoveWorkspace{
		workspaceRepo: wr,
	}
}

func (s *RemoveWorkspace) Handle(ctx context.Context, cmd contract.RemoveWorkspace) (bool, error) {
	if err := s.workspaceRepo.Delete(ctx, cmd.WorkspaceID); err != nil {
		return false, err
	}
	return true, nil
}
