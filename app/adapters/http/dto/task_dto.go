package dto

import (
	"time"

	"gitlab.com/shaninalex/lumna/app/models"
)

type BoardTaskDto struct {
	BoardId  int `json:"board_id"`
	ColumnId int `json:"column_id"`
	Position int `json:"position"`
}

func ToBoardTaskDto(bt models.TaskBoard) BoardTaskDto {
	return BoardTaskDto{
		BoardId:  bt.BoardId,
		ColumnId: bt.ColumnId,
		Position: bt.Position,
	}
}

func ToBoardTaskDtoList(bts []models.TaskBoard) []BoardTaskDto {
	result := []BoardTaskDto{}
	for _, bt := range bts {
		result = append(result, ToBoardTaskDto(bt))
	}
	return result
}

type TaskDto struct {
	ID           int            `json:"id"`
	Title        string         `json:"title"`
	Body         string         `json:"body"`
	Completed    bool           `json:"completed"`
	Meta         string         `json:"meta"`
	ProjectId    int            `json:"project_id"`
	Boards       []BoardTaskDto `json:"boards"`
	OwnerId      int            `json:"owner_id"`
	AssigneesIDs []int          `json:"assignees_ids"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`

	TaskEvents []EntityEventDTO `json:"task_events"`
}

type KanbanMoveTaskDto struct {
	BoardId int `json:"board_id"`
	models.RearangeTask
}
