package contract

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/core/bus"
)

type Reader interface {
	DisplayNames(ctx context.Context, ids []int) (map[int]string, error)
	Exists(ctx context.Context, id int) (bool, error)
}

// Authenticator verifies a password without letting the hash leave identity
type Authenticator interface {
	VerifyPassword(ctx context.Context, email string, password bus.Secret) (int, error)
}

// Provisioner resolves an email to an identity id, creating the identity if
// this is the first time we see it — just-in-time provisioning for logins
// through an external provider, where no registration ever happened.
type Provisioner interface {
	EnsureIdentityByEmail(ctx context.Context, email, fullName string) (int, error)
}

type Mailer interface {
	Send(ctx context.Context, to, template string, data map[string]any) error
}
