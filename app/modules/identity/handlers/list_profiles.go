package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/identity/contract"
	"gitlab.com/shaninalex/lumna/app/modules/identity/internal/domain"
)

type ListProfiles struct {
	identities domain.IdentityRepo
}

func NewListProfiles(
	i domain.IdentityRepo,
) *ListProfiles {
	return &ListProfiles{i}
}

func (u *ListProfiles) Handle(ctx context.Context, cmd contract.ListProfiles) (contract.ListProfilesView, error) {
	var zero contract.ListProfilesView
	identities, err := u.identities.List(ctx, cmd.Limit, cmd.Offset)
	if err != nil {
		return zero, err
	}
	profiles := make([]contract.ProfileView, len(identities))
	for i := range identities {
		profiles[i] = contract.ProfileView{
			ID:       identities[i].ID,
			Email:    identities[i].Email,
			FullName: identities[i].FullName,
			Active:   identities[i].Active,
		}
	}
	return contract.ListProfilesView{
		Profiles: profiles,
		Limit:    cmd.Limit,
		Offset:   cmd.Offset,
	}, nil
}
