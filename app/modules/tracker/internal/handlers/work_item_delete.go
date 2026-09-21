package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
)

type WorkItemDelete struct {
	repo domain.WorkingItemRepo
}

func NewWorkItemDelete(repo domain.WorkingItemRepo) *WorkItemDelete {
	return &WorkItemDelete{
		repo: repo,
	}
}

func (s *WorkItemDelete) Handle(ctx context.Context, cmd contract.WorkItemDelete) (bool, error) {
	if err := s.repo.Delete(ctx, cmd.WorkItemId); err != nil {
		return false, err
	}
	return true, nil
}
