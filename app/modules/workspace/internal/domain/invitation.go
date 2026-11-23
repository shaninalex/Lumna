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

func (s *Invitation) Revoke(t time.Time) error {
	if s.IsAccepted() {
		return InvitationAlreadyAcceptedError
	}
	if s.IsRevoked() {
		return InvitationAlreadyRevokedError
	}

	s.RevokedAt = t
	s.TokenHash = ""

	return nil
}

func (s *Invitation) CanAccept(t time.Time) error {
	if s.IsAccepted() {
		return InvitationAlreadyAcceptedError
	}
	if s.IsRevoked() || s.IsExpired(t) {
		return InvitationExpiredOrAcceptedError
	}
	return nil
}

func (s *Invitation) Accept(t time.Time) error {
	if err := s.CanAccept(t); err != nil {
		return err
	}

	s.AcceptedAt = t
	s.TokenHash = ""

	return nil
}

func (s *Invitation) IsExpired(t time.Time) bool {
	return t.After(s.ExpiresAt)
}

func (s *Invitation) IsRevoked() bool {
	return !s.RevokedAt.IsZero()
}

func (s *Invitation) IsAccepted() bool {
	return !s.AcceptedAt.IsZero()
}
