package contract

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/bus"
)

type RefreshTokensList struct {
	IdentityID int
}

func (RefreshTokensList) Permission() (action string, scope int) { return "", 0 }

func AskRefreshTokensList(ctx context.Context, a *core.App, q RefreshTokensList) ([]RefreshTokenView, error) {
	return bus.Ask[RefreshTokensList, []RefreshTokenView](ctx, a.Queries, q)
}
