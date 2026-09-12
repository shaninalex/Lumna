package domain

import (
	"time"

	"gitlab.com/shaninalex/lumna/app/lib/ptr"
)

type Provider string

var (
	EmailCredentialProvider Provider = "email"
)

type Credential struct {
	ID             int
	IdentityID     int
	Provider       string
	ProviderUserID *string
	Email          *string
	PasswordHash   *string
	CreatedAt      time.Time
}

func NewPasswordCredential(identityId int, hash, email string) *Credential {
	return &Credential{
		IdentityID:   identityId,
		Provider:     string(EmailCredentialProvider),
		Email:        ptr.P(email),
		PasswordHash: ptr.P(hash),
	}
}
