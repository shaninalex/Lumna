package handlers

import (
	"context"
	"errors"

	"gitlab.com/shaninalex/lumna/app/modules/auth/contract"
	"gitlab.com/shaninalex/lumna/app/modules/auth/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type Logout struct {
	tokens  domain.Tokens
	refresh domain.RefreshTokenRepo
	clock   clock.Clock
}

func NewLogout(t domain.Tokens, r domain.RefreshTokenRepo, clk clock.Clock) *Logout {
	return &Logout{tokens: t, refresh: r, clock: clk}
}

// Handle — implements contract.Logout
func (u *Logout) Handle(ctx context.Context, cmd contract.Logout) (contract.LogoutView, error) {
	presented := cmd.RefreshToken.Reveal()
	if presented == "" {
		return contract.LogoutView{Revoked: false}, nil
	}

	stored, err := u.refresh.ByHash(ctx, u.tokens.HashRefresh(presented))
	if err != nil {
		if errors.Is(err, domain.ErrTokenUnknown) {
			return contract.LogoutView{Revoked: false}, nil
		}
		return contract.LogoutView{}, err
	}

	if err = stored.Usable(u.clock.Now()); err != nil {
		return contract.LogoutView{Revoked: false}, nil
	}

	stored.Revoke()

	if _, err = u.refresh.Save(ctx, stored); err != nil {
		return contract.LogoutView{}, err
	}

	return contract.LogoutView{Revoked: true}, nil
}
