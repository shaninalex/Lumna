package domain

import "time"

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
