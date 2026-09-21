package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/identity/contract"
	"gitlab.com/shaninalex/lumna/app/modules/identity/internal/domain"
)

// ListProfiles - AskListProfiles query
type ListProfiles struct {
	identities domain.IdentityRepo
}

func NewListProfiles(
	i domain.IdentityRepo,
) *ListProfiles {
	return &ListProfiles{i}
}

func (u *ListProfiles) Handle(ctx context.Context, q contract.ListProfiles) (contract.ListProfilesView, error) {
	var zero contract.ListProfilesView
	identities, err := u.identities.List(ctx, q.Limit, q.Offset)
	if err != nil {
		return zero, err
	}
	profiles := make([]contract.ProfileView, len(identities))
	for i := range identities {
		profiles[i] = toProfileView(identities[i])
	}
	return contract.ListProfilesView{
		Profiles: profiles,
		Limit:    q.Limit,
		Offset:   q.Offset,
		// add total?
	}, nil
}
