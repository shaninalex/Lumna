package handlers

import (
	"context"
	"time"

	"gitlab.com/shaninalex/lumna/app/modules/auth/contract"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type SessionRefresher interface {
	Refresh(ctx context.Context, plain string, now time.Time) (contract.SessionView, error)
}

type RefreshSession struct {
	refresher SessionRefresher
	clock     clock.Clock
}

func NewRefreshSession(r SessionRefresher, clk clock.Clock) *RefreshSession {
	return &RefreshSession{
		refresher: r,
		clock:     clk,
	}
}

// Handle — implements contract.RefreshSession.
func (s *RefreshSession) Handle(ctx context.Context, cmd contract.RefreshSession) (contract.SessionView, error) {
	return s.refresher.Refresh(ctx, cmd.RefreshToken.Reveal(), s.clock.Now())
}
