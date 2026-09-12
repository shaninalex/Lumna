package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type AddIdentityToWorkspace struct {
	repo  domain.IdentityWorkspaceRepo
	clock clock.Clock
}

func NewAddIdentityToWorkspace(repo domain.IdentityWorkspaceRepo, clk clock.Clock) *AddIdentityToWorkspace {
	return &AddIdentityToWorkspace{
		repo:  repo,
		clock: clk,
	}
}

func (s *AddIdentityToWorkspace) Handle(ctx context.Context, cmd contract.AddIdentityToWorkspace) (bool, error) {
	err := s.repo.Save(ctx, &domain.IdentityWorkspace{
		IdentityID:  cmd.IdentityId,
		WorkspaceID: cmd.WorkspaceId,
		CreatedAt:   s.clock.Now(),
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
