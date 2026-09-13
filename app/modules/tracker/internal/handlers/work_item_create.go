package handlers

import (
	"context"
	"time"

	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
)

type WorkItemCreate struct{}

func NewWorkItemCreate() *WorkItemCreate {
	return &WorkItemCreate{}
}

func (s *WorkItemCreate) Handle(ctx context.Context, cmd contract.WorkItemCreate) (contract.WorkItemView, error) {
	return contract.WorkItemView{
		Id:          0,
		Title:       cmd.Title,
		Description: cmd.Description,
		ProjectId:   cmd.ProjectId,
		StageId:     cmd.StageId,
		ScopeId:     cmd.ScopeId,
		Rank:        0.0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}
