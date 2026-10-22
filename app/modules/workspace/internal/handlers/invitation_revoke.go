package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type InvitationRevoke struct {
	repo  domain.InvitationRepo
	clock clock.Clock
}

func NewInvitationRevoke(repo domain.InvitationRepo, clock clock.Clock) *InvitationRevoke {
	return &InvitationRevoke{repo: repo, clock: clock}
}

func (s *InvitationRevoke) Handle(ctx context.Context, cmd contract.RevokeInvitation) (contract.InvitationView, error) {
	invitation, err := s.repo.GetById(ctx, cmd.InvitationId)
	if err != nil {
		return contract.InvitationView{}, err
	}
	invitation.Revoke(s.clock.Now())
	if err = s.repo.Update(ctx, invitation); err != nil {
		return contract.InvitationView{}, err
	}
	return toInvitationView(invitation), nil
}
