package handlers

import (
	"context"
	"time"

	"gitlab.com/shaninalex/lumna/app/modules/auth/contract"
	"gitlab.com/shaninalex/lumna/app/modules/auth/internal/domain"
)

type RefreshSession struct {
	tokens  domain.Tokens
	refresh domain.RefreshTokenRepo
}

func NewRefreshSession(t domain.Tokens, r domain.RefreshTokenRepo) *RefreshSession {
	return &RefreshSession{tokens: t, refresh: r}
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

	now := time.Now()
	if err := stored.Usable(now); err != nil {
		return zero, err
	}

	stored.Revoke()
	if err := u.refresh.Save(ctx, stored); err != nil {
		return zero, err
	}

	return issueSession(ctx, u.tokens, u.refresh, stored.IdentityID, now)
}
