package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type InvitationValidate struct {
	repo   domain.InvitationRepo
	clock  clock.Clock
	hasher domain.TokenHasher
}

func NewInvitationValidate(repo domain.InvitationRepo, hasher domain.TokenHasher, clk clock.Clock) *InvitationValidate {
	return &InvitationValidate{
		repo:   repo,
		clock:  clk,
		hasher: hasher,
	}
}

func (s *InvitationValidate) Handle(ctx context.Context, q contract.InvitationValidate) (contract.InvitationView, error) {
	invitation, err := s.repo.GetByHash(ctx, s.hasher.Hash(q.Hash))
	if err != nil {
		return contract.InvitationView{}, err
	}

	if err = invitation.CanAccept(s.clock.Now()); err != nil {
		return contract.InvitationView{}, domain.InvitationExpiredOrAcceptedError
	}

	return toInvitationView(invitation), nil
}
