package handlers

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/auth/contract"
	"gitlab.com/shaninalex/lumna/app/modules/auth/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

// InvitationSession issues an invitation token. There is no refresh half:
// when it expires, the user opens the invitation link again.
type InvitationSession struct {
	tokens domain.TokenCodec[contract.InvitationClaims]
	clock  clock.Clock
}

func NewInvitationSession(t domain.TokenCodec[contract.InvitationClaims], clk clock.Clock) *InvitationSession {
	return &InvitationSession{
		tokens: t,
		clock:  clk,
	}
}

// Handle — implements contract.InvitationSession.
func (s *InvitationSession) Handle(_ context.Context, cmd contract.InvitationSession) (contract.SessionView, error) {
	token, ttl, err := s.tokens.Issue(contract.InvitationClaims{
		Email:       cmd.Email,
		Role:        cmd.Role,
		WorkspaceId: cmd.WorkspaceId,
	}, s.clock.Now())
	if err != nil {
		return contract.SessionView{}, err
	}
	return contract.SessionView{
		Email:               cmd.Email,
		Role:                cmd.Role,
		AccessToken:         token,
		AccessTokenDuration: ttl,
	}, nil
}
