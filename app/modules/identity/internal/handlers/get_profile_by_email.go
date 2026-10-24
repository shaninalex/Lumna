package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/identity/contract"
	"gitlab.com/shaninalex/lumna/app/modules/identity/internal/domain"
)

// GetProfileByEmail - contract.AskGetProfileByEmail query
type GetProfileByEmail struct {
	repo domain.IdentityRepo
}

func NewGetProfileByEmail(
	repo domain.IdentityRepo,
) *GetProfileByEmail {
	return &GetProfileByEmail{repo}
}

func (s *GetProfileByEmail) Handle(ctx context.Context, q contract.GetProfileByEmail) (contract.ProfileView, error) {
	identity, err := s.repo.ByEmail(ctx, q.Email)
	if err != nil {
		return contract.ProfileView{}, err
	}
	return toProfileView(identity), nil
}
