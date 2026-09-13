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

func (s *WorkspaceList) Handle(ctx context.Context, q contract.WorkspaceList) ([]contract.WorkspaceView, error) {
	result, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	workspaces := make([]contract.WorkspaceView, len(result))
	for i, w := range result {
		workspaces[i] = contract.WorkspaceView{
			Id:         w.ID,
			Title:      w.Title,
			OwnerEmail: w.OwnerEmail,
			Active:     w.Active,
			CreatedAt:  w.CreatedAt,
			UpdatedAt:  w.UpdatedAt,
		}
	}

	return workspaces, nil
}
