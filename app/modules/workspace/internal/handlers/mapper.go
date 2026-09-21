package handlers

import (
	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
)

func toProjectView(project domain.Project) contract.ProjectView {
	return contract.ProjectView{
		Id:          project.ID,
		Title:       project.Title,
		WorkspaceId: project.WorkspaceId,
		OwnerId:     project.OwnerId,
		Meta:        project.Meta,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
	}
}

func toWorkspaceView(workspace domain.Workspace) contract.WorkspaceView {
	return contract.WorkspaceView{
		Id:         workspace.ID,
		Title:      workspace.Title,
		OwnerEmail: workspace.OwnerEmail,
		Active:     workspace.Active,
		CreatedAt:  workspace.CreatedAt,
		UpdatedAt:  workspace.UpdatedAt,
	}
}
