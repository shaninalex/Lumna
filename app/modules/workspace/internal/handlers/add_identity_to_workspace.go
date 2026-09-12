package handlers

import (
	"context"
	"time"

	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
)

type AddIdentityToWorkspace struct {
	repo domain.IdentityWorkspaceRepo
}

func NewAddIdentityToWorkspace(repo domain.IdentityWorkspaceRepo) *AddIdentityToWorkspace {
	return &AddIdentityToWorkspace{
		repo: repo,
	}
}

func (s *AddIdentityToWorkspace) Handle(ctx context.Context, cmd contract.AddIdentityToWorkspace) (bool, error) {
	err := s.repo.Save(ctx, &domain.IdentityWorkspace{
		IdentityID:  cmd.IdentityId,
		WorkspaceID: cmd.WorkspaceId,
		CreatedAt:   time.Now(),
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
