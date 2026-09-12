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
	tokens     domain.Tokens
	refresh    domain.RefreshTokenRepo
	clock      clock.Clock
}

func NewEmailLogin(
	i contract.IdentityAuthenticator,
	t domain.Tokens,
	r domain.RefreshTokenRepo,
	clk clock.Clock,
) *EmailLogin {
	return &EmailLogin{identities: i, tokens: t, refresh: r, clock: clk}
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

	return issueSession(ctx, u.tokens, u.refresh, identityID, u.clock.Now())
}

// issueSession mints an access/refresh pair and stores the refresh half.
// Shared by login and rotation so both can never drift apart.
func issueSession(
	ctx context.Context,
	tokens domain.Tokens,
	repo domain.RefreshTokenRepo,
	identityID int,
	now time.Time,
) (contract.SessionView, error) {
	var zero contract.SessionView

	access, accessTTL, err := tokens.IssueAccess(identityID, now)
	if err != nil {
		return zero, err
	}

	plain, hash, refreshTTL, err := tokens.IssueRefresh()
	if err != nil {
		return zero, err
	}

	// Transaction is already open (command middleware) — nothing to manage here.
	if err := repo.Save(ctx, domain.NewRefreshToken(identityID, hash, refreshTTL, now)); err != nil {
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
