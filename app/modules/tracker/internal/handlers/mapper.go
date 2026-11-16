package handlers

import (
	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
)

func toWorkItemView(w domain.WorkItem) contract.WorkItemView {
	return contract.WorkItemView{
		Id:          w.ID,
		Title:       w.Title,
		Description: w.Description,
		ProjectId:   w.ProjectID,
		StageId:     w.StageID,
		ScopeId:     w.ScopeID,
		Rank:        w.Rank,
		DueTo:       w.DueTo,
		CreatedAt:   w.CreatedAt,
		UpdatedAt:   w.UpdatedAt,
		Assignees:   w.AssigneeIDs,
	}
}

func toWorkItemViews(items []domain.WorkItem) []contract.WorkItemView {
	views := make([]contract.WorkItemView, len(items))
	for i, item := range items {
		views[i] = toWorkItemView(item)
	}
	return views
}

func toScopeView(scope domain.Scope) contract.ScopeView {
	return contract.ScopeView{
		Id:          scope.ID,
		Name:        scope.Name,
		Description: scope.Description,
		ProjectId:   scope.ProjectID,
		CreatedAt:   scope.CreatedAt,
		UpdatedAt:   scope.UpdatedAt,
	}
}

func toScopeViews(scopes []domain.Scope) []contract.ScopeView {
	views := make([]contract.ScopeView, len(scopes))
	for i, scope := range scopes {
		views[i] = toScopeView(scope)
	}
	return views
}

func toStageView(stage domain.Stage) contract.StageView {
	return contract.StageView{
		Id:          stage.ID,
		ScopeId:     stage.ScopeID,
		Name:        stage.Name,
		Description: stage.Description,
		Category:    string(stage.Category),
		Position:    stage.Position,
		WIPLimit:    stage.WIPLimit,
		CreatedAt:   stage.CreatedAt,
		UpdatedAt:   stage.UpdatedAt,
	}
}

func toStageViews(stages []domain.Stage) []contract.StageView {
	views := make([]contract.StageView, len(stages))
	for i, stage := range stages {
		views[i] = toStageView(stage)
	}
	return views
}
