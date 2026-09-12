package contract

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/bus"
)

// EmailLogin — authenticate by email and password, issue a session.
type EmailLogin struct {
	Email    string
	Password bus.Secret
}

func (EmailLogin) Permission() (action string, scope int) { return "", 0 } // public

func ExecEmailLogin(ctx context.Context, a *core.App, c EmailLogin) (SessionView, error) {
	return bus.Execute[EmailLogin, SessionView](ctx, a.Commands, c)
}

// RefreshSession — exchange a refresh token for a new pair, rotating the old one.
type RefreshSession struct {
	RefreshToken bus.Secret
}

func (RefreshSession) Permission() (action string, scope int) { return "", 0 } // the token IS the credential

func ExecRefreshSession(ctx context.Context, a *core.App, c RefreshSession) (SessionView, error) {
	return bus.Execute[RefreshSession, SessionView](ctx, a.Commands, c)
}

// Logout — revoke a single refresh token.
type Logout struct {
	RefreshToken bus.Secret
}

func (Logout) Permission() (action string, scope int) { return "", 0 }

func ExecLogout(ctx context.Context, a *core.App, c Logout) (LogoutView, error) {
	return bus.Execute[Logout, LogoutView](ctx, a.Commands, c)
}
