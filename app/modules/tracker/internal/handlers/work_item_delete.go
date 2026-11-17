package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/core/bus"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
)

type WorkItemDelete struct {
	repo domain.WorkingItemRepo
	bus  *bus.EventBus
}

func NewWorkItemDelete(repo domain.WorkingItemRepo, bus *bus.EventBus) *WorkItemDelete {
	return &WorkItemDelete{
		repo: repo,
		bus:  bus,
	}
}

func (s *WorkItemDelete) Handle(ctx context.Context, cmd contract.WorkItemDelete) (bool, error) {
	if err := s.repo.Delete(ctx, cmd.WorkItemId); err != nil {
		return false, err
	}
	_ = s.bus.Publish(ctx, contract.WorkItemChanged{
		WorkItemId: cmd.WorkItemId,
		ActorId:    actorID(ctx),
	})
	return true, nil
}
