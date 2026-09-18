package task

import (
	"encoding/json"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
)

type taskDTO struct {
	ID           int        `json:"id"`
	Title        string     `json:"title"`
	Body         *string    `json:"body"`
	Completed    bool       `json:"completed"`
	Meta         string     `json:"meta"`
	ProjectId    int        `json:"project_id"`
	BoardId      *int       `json:"board_id"`
	ColumnId     *int       `json:"column_id"`
	Position     float64    `json:"position"`
	OwnerId      int        `json:"owner_id"`
	AssigneesIDs []int      `json:"assignees_ids"`
	DueTo        *time.Time `json:"due_to"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

type BoardTaskDto struct {
	BoardId  int     `json:"board_id"`
	ColumnId int     `json:"column_id"`
	Position float64 `json:"position"`
}

func toTaskDTO(w contract.WorkItemView) taskDTO {
	task := taskDTO{
		ID:           w.Id,
		Title:        w.Title,
		Body:         w.Description,
		Completed:    false,
		Meta:         "",
		ProjectId:    w.ProjectId,
		BoardId:      w.ScopeId,
		ColumnId:     w.StageId,
		Position:     w.Rank,
		OwnerId:      0,
		AssigneesIDs: []int{},
		DueTo:        w.DueTo,
		CreatedAt:    w.CreatedAt,
		UpdatedAt:    w.UpdatedAt,
	}

	return task
}

type taskCreateDTO struct {
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	ProjectId int        `json:"project_id"`
	Position  float64    `json:"position"`
	ColumnId  int        `json:"column_id"`
	BoardId   int        `json:"board_id"`
	DueTo     *time.Time `json:"due_to"`
}

func (a taskCreateDTO) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Title, validation.Required),
	)
}

type columnDTO struct {
	Id        int        `json:"id"`
	Title     string     `json:"title"`
	Meta      columnMeta `json:"meta"`
	BoardId   int        `json:"board_id"`
	Position  float64    `json:"position"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type columnMeta struct {
	Color    string `json:"color"`
	Icon     string `json:"icon"`
	Expanded bool   `json:"expanded"`
}

func toColumnDTO(s contract.StageView) columnDTO {
	return columnDTO{
		Id:      s.Id,
		Title:   s.Name,
		BoardId: s.ScopeId,
		Meta: columnMeta{
			Color:    "default",
			Icon:     "default",
			Expanded: true,
		},
		Position:  s.Position,
		CreatedAt: s.CreatedAt,
		UpdatedAt: &s.UpdatedAt,
	}
}

type boardAction string

var (
	boardActionMoveColumn  boardAction = "move_column"
	boardActionMoveTask    boardAction = "move_task"
	boardActionChangeStage boardAction = "change_stage"
)

type boardActionMoveTaskDTO struct {
	BoardId  int     `json:"board_id"`
	TaskId   int     `json:"task_id"`
	Position float64 `json:"position"`
}

type boardActionMoveColumnDTO struct {
	BoardId  int     `json:"board_id"`
	ColumnId int     `json:"column_id"`
	Position float64 `json:"position"`
}

type boardActionTransferTaskDTO struct {
	TaskId   int     `json:"task_id"`
	BoardId  int     `json:"board_id"`
	ColumnId int     `json:"column_id"`
	Position float64 `json:"position"`
}

type boardActionPayload struct {
	Action boardAction     `json:"action"`
	Data   json.RawMessage `json:"data"`
}

type taskUpdateDTO struct {
	TaskId int    `json:"task_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (a taskUpdateDTO) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.TaskId, validation.Required),
		validation.Field(&a.TaskId, validation.Min(1)), // task id should not be 0
		validation.Field(&a.Title, validation.Required),
	)
}
