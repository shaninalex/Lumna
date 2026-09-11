package domain

import "time"

// RefreshToken is the stored half of a session. The plaintext token is handed
// to the client once and never persisted — only its hash lives here, so a
// leaked database does not hand out sessions.
type RefreshToken struct {
	ID         int
	IdentityID int
	Hash       string
	ClientID   *string
	Scopes     *string
	ExpiresAt  time.Time
	Revoked    bool
	CreatedAt  time.Time
}

func NewRefreshToken(identityID int, hash string, ttl time.Duration, now time.Time) *RefreshToken {
	return &RefreshToken{
		IdentityID: identityID,
		Hash:       hash,
		ExpiresAt:  now.Add(ttl),
		Revoked:    false,
		CreatedAt:  now,
	}
}

// Usable holds the invariant of a session: alive and not past its end.
func (t *RefreshToken) Usable(now time.Time) error {
	if t.Revoked {
		return ErrTokenRevoked
	}
	if !now.Before(t.ExpiresAt) {
		return ErrTokenExpired
	}
	return nil
}

func (t *RefreshToken) Revoke() { t.Revoked = true }
