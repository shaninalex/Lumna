package domain

import (
	"time"
)

type Identity struct {
	ID       int
	Email    string
	FullName string
	Active   bool
	Created  time.Time
	Updated  time.Time
}

type Hasher interface {
	Hash(plain string) (string, error)
	Verify(hash, plain string) error
}

// NewIdentity - makes new identity
func NewIdentity(email, fullName string, now time.Time) Identity {
	return Identity{
		Email:    email,
		FullName: fullName,
		Active:   true,
		Created:  now,
	}
}
