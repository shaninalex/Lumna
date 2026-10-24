package contract

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/bus"
	"gitlab.com/shaninalex/lumna/app/core/secure"
)

// EmailLogin — authenticate by email and password, issue a session.
type EmailLogin struct {
	Email    string
	Password secure.Secret
}

func (EmailLogin) Permission() (action string, scope int) { return "", 0 } // public

func ExecEmailLogin(ctx context.Context, a *core.App, c EmailLogin) (SessionView, error) {
	return bus.Execute[EmailLogin, SessionView](ctx, a.Commands, c)
}

// RefreshSession — exchange a refresh token for a new pair, rotating the old one.
type RefreshSession struct {
	RefreshToken secure.Secret
}

func (RefreshSession) Permission() (action string, scope int) { return "", 0 } // the token IS the credential

func ExecRefreshSession(ctx context.Context, a *core.App, c RefreshSession) (SessionView, error) {
	return bus.Execute[RefreshSession, SessionView](ctx, a.Commands, c)
}

// Logout — revoke a single refresh token.
type Logout struct {
	RefreshToken secure.Secret
}

func (Logout) Permission() (action string, scope int) { return "", 0 }

func ExecLogout(ctx context.Context, a *core.App, c Logout) (LogoutView, error) {
	return bus.Execute[Logout, LogoutView](ctx, a.Commands, c)
}

type InvitationSession struct {
	Email        string
	Role         string
	InvitationID int
	WorkspaceId  int
}

func (InvitationSession) Permission() (action string, scope int) { return "", 0 }

func ExecInvitationSession(ctx context.Context, a *core.App, cmd InvitationSession) (SessionView, error) {
	return bus.Execute[InvitationSession, SessionView](ctx, a.Commands, cmd)
}
