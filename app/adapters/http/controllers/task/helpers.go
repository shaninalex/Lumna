package task

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/api/adapters"
	"gitlab.com/shaninalex/lumna/app/api/dto"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/services"
)

func (s *TaskController) queryTaskList(ctx context.Context, q services.ServiceTaskListQuery) ([]dto.TaskDto, error) {
	tasks, err := s.taskService.List(ctx, q)
	if err != nil {
		return nil, err
	}

	ids := make([]int, len(tasks))
	for i, t := range tasks {
		ids[i] = t.ID
	}

	events := []models.EntityEvent{}
	events, err = s.entityEventService.ListByEntityIds(ctx, ids, "task")
	if err != nil {
		return nil, err
	}
	return adapters.ToTaskDtoList(tasks, events), nil
}
