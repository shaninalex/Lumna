package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
)

type RemoveProject struct {
	repo domain.ProjectRepo
}

func NewRemoveProject(repo domain.ProjectRepo) *RemoveProject {
	return &RemoveProject{repo}
}

func (s *RemoveProject) Handle(ctx context.Context, cmd contract.RemoveProject) (bool, error) {
	if err := s.repo.Remove(ctx, cmd.ProjectID); err != nil {
		return false, err
	}

	return true, nil
}
