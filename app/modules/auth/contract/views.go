package contract

import "time"

// SessionView is what a successful authentication produces: tokens, and the
// id they belong to.
type SessionView struct {
	Email                string
	Role                 string
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

type RefreshTokenView struct {
	ID         int
	IdentityID int
	ClientID   string
	Scopes     string
	ExpiresAt  time.Time
	Revoked    bool
	CreatedAt  time.Time
	//Hash       string
}
