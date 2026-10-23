package handlers

import (
	"context"
	"time"

	"gitlab.com/shaninalex/lumna/app/modules/auth/contract"
	"gitlab.com/shaninalex/lumna/app/modules/auth/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type EmailLogin struct {
	identities contract.IdentityAuthenticator // port: implemented by another module
	sessions   *SessionIssuer
	clock      clock.Clock
}

func NewEmailLogin(i contract.IdentityAuthenticator, s *SessionIssuer, clk clock.Clock) *EmailLogin {
	return &EmailLogin{identities: i, sessions: s, clock: clk}
}

// Handle — implements bus.RegisterCommand.
func (u *EmailLogin) Handle(ctx context.Context, cmd contract.EmailLogin) (contract.SessionView, error) {
	var zero contract.SessionView

	// The whole cross-module step: an email and a password go out, an id comes
	// back. No hash, no profile, no type from the other module's contract.
	identityID, err := u.identities.VerifyPassword(ctx, cmd.Email, cmd.Password)
	if err != nil {
		return zero, err
	}

	return u.sessions.Issue(ctx, identityID, u.clock.Now())
}

// SessionIssuer mints an access/refresh pair and stores the refresh half.
// Shared by login and rotation so both can never drift apart.
type SessionIssuer struct {
	access  domain.TokenCodec[domain.AccessClaims]
	refresh domain.RefreshTokens
	repo    domain.RefreshTokenRepo
}

func NewSessionIssuer(
	a domain.TokenCodec[domain.AccessClaims],
	r domain.RefreshTokens,
	repo domain.RefreshTokenRepo,
) *SessionIssuer {
	return &SessionIssuer{access: a, refresh: r, repo: repo}
}

func (s *SessionIssuer) Issue(ctx context.Context, identityID int, now time.Time) (contract.SessionView, error) {
	var zero contract.SessionView

	access, accessTTL, err := s.access.Issue(domain.AccessClaims{IdentityID: identityID}, now)
	if err != nil {
		return zero, err
	}

	plain, hash, refreshTTL, err := s.refresh.Issue()
	if err != nil {
		return zero, err
	}

	if _, err = s.repo.Save(ctx, domain.NewRefreshToken(identityID, hash, refreshTTL, now)); err != nil {
		return zero, err
	}

	return contract.SessionView{
		IdentityID:           identityID,
		AccessToken:          access,
		AccessTokenDuration:  accessTTL,
		RefreshToken:         plain,
		RefreshTokenDuration: refreshTTL,
	}, nil
}
