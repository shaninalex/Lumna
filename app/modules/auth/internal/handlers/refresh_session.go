package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/auth/contract"
	"gitlab.com/shaninalex/lumna/app/modules/auth/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type RefreshSession struct {
	tokens  domain.Tokens
	refresh domain.RefreshTokenRepo
	clock   clock.Clock
}

func NewRefreshSession(t domain.Tokens, r domain.RefreshTokenRepo, clk clock.Clock) *RefreshSession {
	return &RefreshSession{tokens: t, refresh: r, clock: clk}
}

// Handle — implements contract.RefreshSession.
func (u *RefreshSession) Handle(ctx context.Context, cmd contract.RefreshSession) (contract.SessionView, error) {
	var zero contract.SessionView

	presented := cmd.RefreshToken.Reveal()
	if presented == "" {
		return zero, domain.ErrEmptyToken
	}

	stored, err := u.refresh.ByHash(ctx, u.tokens.HashRefresh(presented))
	if err != nil {
		return zero, err
	}

	now := u.clock.Now()
	if err := stored.Usable(now); err != nil {
		return zero, err
	}

	stored.Revoke()
	if err := u.refresh.Save(ctx, stored); err != nil {
		return zero, err
	}

	return issueSession(ctx, u.tokens, u.refresh, stored.IdentityID, now)
}
