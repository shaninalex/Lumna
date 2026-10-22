package domain

import (
	"time"
)

type Invitation struct {
	ID          int
	WorkspaceID int
	Email       string
	Role        string
	TokenHash   string
	InvitedBy   int
	ExpiresAt   time.Time
	AcceptedAt  time.Time
	RevokedAt   time.Time
	CreatedAt   time.Time
}

// Revoke - revoke invitation. Also removes TokenHash
func (s *Invitation) Revoke(t time.Time) {
	s.RevokedAt = t
	s.TokenHash = ""
}
