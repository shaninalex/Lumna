package contract

import "context"

type GetProfile struct{ IdentityID int }

func (q GetProfile) Permission() (string, int) { return "identity.read", q.IdentityID }

func AskGetProfile(ctx context.Context, a *core.App, q GetProfile) (ProfileView, error) {
	return bus.Ask[GetProfile, ProfileView](ctx, a.Queries, q)
}
