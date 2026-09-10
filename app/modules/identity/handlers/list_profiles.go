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

// Handle — implements bus.RegisterCommand.
func (u *ListProfiles) Handle(ctx context.Context, cmd contract.Register) (contract.ListProfilesView, error) {
	//var zero contract.ListProfilesView

	return contract.ListProfilesView{}, nil
}
