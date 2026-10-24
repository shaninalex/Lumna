package contract

import (
	"context"
	"errors"
	"net/mail"

	"gitlab.com/shaninalex/lumna/app/core/secure"
)

type IdentityAuthenticator interface {

	// VerifyPassword authenticate user by email/password
	VerifyPassword(ctx context.Context, email string, password secure.Secret) (int, error)
}

type IdentityProvisioner interface {

	// EnsureIdentityByEmail - Get or create identity. Need for OAuth flow.
	EnsureIdentityByEmail(ctx context.Context, email, fullName string) (int, error)
}

type Verifier[T any] interface {
	Verify(ctx context.Context, token string) (T, error)
}

type InvitationClaims struct {
	Email        string `json:"email"`
	InvitationID int    `json:"invitation_id"`
	Role         string `json:"role"`
	WorkspaceId  int    `json:"workspace_id"`
}

var errInvalidInvitationClaims = errors.New("invalid invitation claims")

func (c InvitationClaims) Validate() error {
	if _, err := mail.ParseAddress(c.Email); err != nil {
		return errors.Join(errInvalidInvitationClaims, err)
	}
	if c.InvitationID == 0 {
		return errInvalidInvitationClaims
	}
	if c.WorkspaceId == 0 {
		return errInvalidInvitationClaims
	}
	return nil
}

type PasswordVerifier interface {
	Verify(pwd secure.Secret) int
}
