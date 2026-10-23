package contract

import (
	"context"
	"errors"
	"net/mail"

	"gitlab.com/shaninalex/lumna/app/core/bus"
)

type IdentityAuthenticator interface {

	// VerifyPassword authenticate user by email/password
	VerifyPassword(ctx context.Context, email string, password bus.Secret) (int, error)
}

type IdentityProvisioner interface {

	// EnsureIdentityByEmail - Get or create identity. Need for OAuth flow.
	EnsureIdentityByEmail(ctx context.Context, email, fullName string) (int, error)
}

type Verifier[T any] interface {
	Verify(ctx context.Context, token string) (T, error)
}

type InvitationClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

var errInvalidInvitationClaims = errors.New("invalid invitation claims")

func (c InvitationClaims) Validate() error {
	if _, err := mail.ParseAddress(c.Email); err != nil {
		return errors.Join(errInvalidInvitationClaims, err)
	}
	if c.Role == "" {
		return errInvalidInvitationClaims
	}
	return nil
}
