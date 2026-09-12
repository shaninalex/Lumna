package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
)

type CreateProject struct {
	repo domain.ProjectRepo
}

func NewCreateProject(repo domain.ProjectRepo) *CreateProject {
	return &CreateProject{
		repo: repo,
	}
}

func (s *CreateProject) Handle(ctx context.Context, cmd contract.CreateProject) (contract.ProjectView, error) {
	project := domain.NewProject(cmd.Title, cmd.WorkspaceId, cmd.OwnerId)
	if err := s.repo.Save(ctx, project); err != nil {
		return contract.ProjectView{}, err
	}

	return contract.ProjectView{
		Id:          project.ID,
		Title:       project.Title,
		WorkspaceId: project.WorkspaceId,
		OwnerId:     project.OwnerId,
		Meta:        project.Meta,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
	}, nil
}
