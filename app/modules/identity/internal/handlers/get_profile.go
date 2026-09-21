package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/identity/contract"
	"gitlab.com/shaninalex/lumna/app/modules/identity/internal/domain"
)

// GetProfile - contract.AskGetProfile query
type GetProfile struct {
	identities domain.IdentityRepo
}

func NewGetProfile(
	i domain.IdentityRepo,
) *GetProfile {
	return &GetProfile{i}
}

func (s *GetProfile) Handle(ctx context.Context, q contract.GetProfile) (contract.ProfileView, error) {
	identity, err := s.identities.ByID(ctx, q.IdentityID)
	if err != nil {
		return contract.ProfileView{}, err
	}
	return toProfileView(identity), nil
}
