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

func (s *InvitationValidate) Handle(ctx context.Context, q contract.InvitationValidate) (bool, error) {
	// The invitation exists.
	invitation, err := s.repo.GetByHash(ctx, s.hasher.q.Hash)
	if err != nil {
		return false, err
	}

	// It has not expired.
	// It has not already been accepted.
	// It has not been revoked.
	if invitation.IsExpired(s.clock.Now()) || invitation.IsRevoked() || invitation.IsAccepted() {
		return false, domain.InvitationExpiredOrAcceptedError
	}

	return true, nil
}
