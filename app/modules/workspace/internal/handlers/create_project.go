package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type CreateProject struct {
	repo  domain.ProjectRepo
	clock clock.Clock
}

func NewCreateProject(repo domain.ProjectRepo, clk clock.Clock) *CreateProject {
	return &CreateProject{
		repo:  repo,
		clock: clk,
	}
}

func (s *CreateProject) Handle(ctx context.Context, cmd contract.CreateProject) (contract.ProjectView, error) {
	newProject, err := s.repo.Save(ctx, domain.NewProject(cmd.Title, cmd.WorkspaceId, cmd.OwnerId, s.clock.Now()))
	if err != nil {
		return contract.ProjectView{}, err
	}
	return toProjectView(newProject), nil
}
