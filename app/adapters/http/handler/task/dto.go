package task

import (
	"time"

	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
)

type taskDTO struct {
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
}

type BoardTaskDto struct {
	BoardId  int     `json:"board_id"`
	ColumnId int     `json:"column_id"`
	Position float64 `json:"position"`
}

func toDTO(w contract.WorkItemView) taskDTO {
	return taskDTO{
		ID:        w.Id,
		Title:     w.Title,
		Body:      w.Description,
		Completed: false,
		Meta:      "",
		ProjectId: w.ProjectId,
		Boards: []BoardTaskDto{
			{
				BoardId:  w.Id,
				ColumnId: w.Id,
				Position: w.Rank,
			},
		},
		OwnerId:      0,
		AssigneesIDs: []int{},
		CreatedAt:    w.CreatedAt,
		UpdatedAt:    w.UpdatedAt,
	}
}

type taskCreateDTO struct {
	Title     string `json:"title"` // required
	Body      string `json:"body"`
	ProjectId int    `json:"project_id"`
	Position  int    `json:"position"` // > 0
	ColumnId  int    `json:"column_id"`
	BoardId   int    `json:"board_id"`
}
