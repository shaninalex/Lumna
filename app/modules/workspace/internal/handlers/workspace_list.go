package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
)

type WorkspaceList struct {
	repo domain.WorkspaceRepo
}

func NewWorkspaceList(repo domain.WorkspaceRepo) *WorkspaceList {
	return &WorkspaceList{repo: repo}
}

func (s *WorkspaceList) Handle(ctx context.Context, _ contract.WorkspaceList) ([]contract.WorkspaceView, error) {
	result, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	workspaces := make([]contract.WorkspaceView, len(result))
	for i, workspace := range result {
		workspaces[i] = toWorkspaceView(workspace)
	}
	return workspaces, nil
}
