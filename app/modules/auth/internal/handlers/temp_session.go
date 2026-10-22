package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/auth/contract"
	"gitlab.com/shaninalex/lumna/app/modules/auth/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type TempSession struct {
	tokens domain.Tokens
	clock  clock.Clock
}

func NewTempSession(tokens domain.Tokens, clk clock.Clock) *TempSession {
	return &TempSession{
		tokens: tokens,
		clock:  clk,
	}
}

func (s *TempSession) Handle(_ context.Context, cmd contract.TempSession) (contract.SessionView, error) {
	token, ttl, err := s.tokens.IssueTempAccess(cmd.Email, cmd.Email, s.clock.Now())
	if err != nil {
		return contract.SessionView{}, err
	}
	return contract.SessionView{
		AccessToken:         token,
		AccessTokenDuration: ttl,
	}, nil
}
