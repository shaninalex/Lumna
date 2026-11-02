package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
)

type RemoveMember struct {
	identityWorkspaceRepo domain.IdentityWorkspaceRepo
}

func NewRemoveMember(identityWorkspaceRepo domain.IdentityWorkspaceRepo) *RemoveMember {
	return &RemoveMember{
		identityWorkspaceRepo: identityWorkspaceRepo,
	}
}

func (s *RemoveMember) Handle(ctx context.Context, cmd contract.RemoveMember) (bool, error) {
	if err := s.identityWorkspaceRepo.Delete(ctx, cmd.IdentityId, cmd.WorkspaceID); err != nil {
		return false, err
	}
	return true, nil
}
