package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
)

type InvitationDelete struct {
	repo domain.InvitationRepo
}

func NewInvitationDelete(repo domain.InvitationRepo) *InvitationDelete {
	return &InvitationDelete{repo: repo}
}

func (s *InvitationDelete) Handle(ctx context.Context, cmd contract.DeleteInvitation) (bool, error) {
	if err := s.repo.Delete(ctx, cmd.InvitationId); err != nil {
		return false, err
	}

	return true, nil
}
