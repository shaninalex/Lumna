package contract

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/bus"
)

type GetProfile struct {
	IdentityID int
}

func (q GetProfile) Permission() (string, int) {
	return "identity.read", q.IdentityID
}

func AskGetProfile(ctx context.Context, a *core.App, q GetProfile) (ProfileView, error) {
	return bus.Ask[GetProfile, ProfileView](ctx, a.Queries, q)
}

type ListProfiles struct {
	Limit  int
	Offset int
}

func (q ListProfiles) Permission() (string, int) {
	return "identity.read", 0
}

func AskListProfiles(ctx context.Context, a *core.App, q ListProfiles) (ListProfilesView, error) {
	return bus.Ask[ListProfiles, ListProfilesView](ctx, a.Queries, q)
}
