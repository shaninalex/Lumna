package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
)

type ProjectList struct {
	repo domain.ProjectRepo
}

func NewProjectList(repo domain.ProjectRepo) *ProjectList {
	return &ProjectList{repo: repo}
}

func (s *ProjectList) Handle(ctx context.Context, list contract.ProjectList) ([]contract.ProjectView, error) {
	result, err := s.repo.ByWorkspaceId(ctx, list.WorkspaceId)
	if err != nil {
		return nil, err
	}
	projects := make([]contract.ProjectView, len(result))
	for i, project := range result {
		projects[i] = contract.ProjectView{
			Id:          project.ID,
			Title:       project.Title,
			WorkspaceId: project.WorkspaceId,
			OwnerId:     project.OwnerId,
			Meta:        project.Meta,
			CreatedAt:   project.CreatedAt,
			UpdatedAt:   project.UpdatedAt,
		}
	}

	return projects, nil
}
