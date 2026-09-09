package adapters

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/dto"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/services"
)

type TaskListQuery struct {
	BoardId *int `form:"board_id,omitempty"`
}

var (
	TaskListParseParamsError = errors.New("no query parameters")
)

func ParseQueryParams(c *gin.Context) (*services.ServiceTaskListQuery, error) {
	q := TaskListQuery{}
	if err := c.ShouldBindQuery(&q); err != nil {
		return nil, err
	}

	if q.BoardId == nil {
		return nil, TaskListParseParamsError
	}

	return &services.ServiceTaskListQuery{
		BoardId: *q.BoardId,
	}, nil
}

func ToTaskDto(task models.Task) dto.TaskDto {
	return dto.TaskDto{
		ID:           task.ID,
		Title:        task.Title,
		Body:         task.Body,
		Completed:    task.Completed,
		Meta:         task.Meta,
		ProjectId:    task.ProjectId,
		Boards:       dto.ToBoardTaskDtoList(task.Boards), // should be single board!
		OwnerId:      task.OwnerId,
		AssigneesIDs: task.AssigneesIDs,
		CreatedAt:    task.CreatedAt,
		UpdatedAt:    task.UpdatedAt,
		TaskEvents:   []dto.EntityEventDTO{},
	}
}

func ToTaskDtoList(tasks []models.Task, events []models.EntityEvent) []dto.TaskDto {
	result := make([]dto.TaskDto, 0, len(tasks))
	for _, task := range tasks {
		td := ToTaskDto(task)

		for _, e := range events {
			if e.EntityId == nil || e.EntityType == nil {
				continue
			}

			if task.ID == *e.EntityId && *e.EntityType == "task" {
				td.TaskEvents = append(td.TaskEvents, dto.ToEntityEventDTO(e))
			}
		}

		result = append(result, td)
	}
	return result
}
