package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/lib"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
)

type StageDelete struct {
	itemsRepo domain.WorkingItemRepo
	stageRepo domain.StageRepo
}

func NewStageDelete(itemsRepo domain.WorkingItemRepo, stageRepo domain.StageRepo) *StageDelete {
	return &StageDelete{
		itemsRepo: itemsRepo,
		stageRepo: stageRepo,
	}
}

func (s *StageDelete) Handle(ctx context.Context, cmd contract.StageDelete) (contract.StageDeleteView, error) {
	response := contract.StageDeleteView{
		StageId: cmd.StageId,
	}

	if cmd.WithTasks {
		stage, err := s.stageRepo.GetById(ctx, cmd.StageId)
		if err != nil {
			return contract.StageDeleteView{}, err
		}
		ids := lib.Map(stage.WorkItems, func(item domain.WorkItem) int {
			return item.ID
		})
		if len(ids) != 0 {
			if _, err = s.itemsRepo.BatchDelete(ctx, ids); err != nil {
				return contract.StageDeleteView{}, err
			}
			response.DeletedTasks = ids
		}
	}

	if _, err := s.stageRepo.Delete(ctx, cmd.StageId); err != nil {
		return contract.StageDeleteView{}, err
	}

	return response, nil
}
