package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type UpdateProject struct {
	repo domain.ProjectRepo
	clk  clock.Clock
}

func NewUpdateProject(repo domain.ProjectRepo, clk clock.Clock) *UpdateProject {
	return &UpdateProject{repo, clk}
}

func (s *UpdateProject) Handle(ctx context.Context, cmd contract.UpdateProject) (contract.ProjectView, error) {
	project, err := s.repo.GetById(ctx, cmd.ProjectID)
	if err != nil {
		return contract.ProjectView{}, err
	}

	project.Title = cmd.Title
	project.UpdatedAt = s.clk.Now()

	newProject, err := s.repo.Update(ctx, project)
	if err != nil {
		return contract.ProjectView{}, err
	}
	return toProjectView(newProject), nil
}
