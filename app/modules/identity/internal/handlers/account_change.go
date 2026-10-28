package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/core/actor"
	"gitlab.com/shaninalex/lumna/app/modules/identity/contract"
	"gitlab.com/shaninalex/lumna/app/modules/identity/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type AccountChange struct {
	identityRepo domain.IdentityRepo
	clk          clock.Clock
}

func NewAccountChange(identityRepo domain.IdentityRepo, clk clock.Clock) *AccountChange {
	return &AccountChange{identityRepo, clk}
}

func (s *AccountChange) Handle(ctx context.Context, cmd contract.AccountChange) (contract.ProfileView, error) {
	identity, err := s.identityRepo.ByID(ctx, cmd.IdentityID)
	if err != nil {
		return contract.ProfileView{}, err
	}

	// account change should be possible ONLY via authorized action by the account itself
	a, ok := actor.From(ctx)
	if !ok || a.IdentityID == 0 || a.IdentityID != identity.ID {
		return contract.ProfileView{}, domain.ErrBadCredentials
	}

	identity.Email = cmd.Email.Address
	identity.FullName = cmd.FullName
	identity.Updated = s.clk.Now()

	if err = s.identityRepo.Update(ctx, identity); err != nil {
		return contract.ProfileView{}, err
	}

	return toProfileView(identity), nil
}
