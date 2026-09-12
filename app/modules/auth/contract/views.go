package contract

import "time"

// SessionView is what a successful authentication produces: tokens, and the
// id they belong to.
type SessionView struct {
	IdentityID           int
	AccessToken          string
	AccessTokenDuration  time.Duration
	RefreshToken         string
	RefreshTokenDuration time.Duration
}

// LogoutView reports whether the presented token was still live when revoked.
type LogoutView struct {
	Revoked bool
}
