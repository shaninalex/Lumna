package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
)

type ActivityList struct {
	repo domain.ActivityRepo
}

func NewActivityList(repo domain.ActivityRepo) *ActivityList {
	return &ActivityList{
		repo: repo,
	}
}

func (s *ActivityList) Handle(ctx context.Context, cmd contract.ActivityList) ([]contract.ActivityView, error) {
	result, err := s.repo.ListById(ctx, cmd.EntityId, cmd.EntityType)
	if err != nil {
		return nil, err
	}
	return toActivityViews(result), nil
}
