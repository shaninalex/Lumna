package board

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
)

type boardDTO struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	ProjectId  int       `json:"project_id"`
	StageCount int       `json:"stage_count"`
	IssueCount int       `json:"issue_count"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at,omitzero"`
}

func toBoardDTO(scope contract.ScopeView) boardDTO {
	return boardDTO{
		Id:         scope.Id,
		Title:      scope.Name,
		ProjectId:  scope.ProjectId,
		StageCount: scope.StageCount,
		IssueCount: scope.IssueCount,
		CreatedAt:  scope.CreatedAt,
		UpdatedAt:  scope.UpdatedAt,
	}
}

type boardCreateDTO struct {
	Title     string `json:"title"`
	ProjectId int    `json:"project_id"`
}

func (a boardCreateDTO) Validate() error {
	return validation.ValidateStruct(
		&a,
		validation.Field(&a.Title, validation.Required, validation.Length(3, 50)),
		validation.Field(&a.ProjectId, validation.Required),
	)
}
