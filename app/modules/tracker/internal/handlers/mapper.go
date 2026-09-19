package handlers

import (
	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
)

func toWorkItemView(w *domain.WorkItem) contract.WorkItemView {
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
		views[i] = toWorkItemView(&item)
	}
	return views
}
