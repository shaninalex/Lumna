// app/modules/identity/internal/domain/identity.go
package domain

import (
	"net/mail"
	"strings"
	"time"
)

type Identity struct {
	ID       int
	Email    string
	FullName string
	Active   bool
	Created  time.Time
}

// Hasher — порт. Argon2 живе в infra, домен знає лише інтерфейс.
type Hasher interface {
	Hash(plain string) (string, error)
	Verify(hash, plain string) error
}

// NewIdentity тримає інваріант «email нормалізований і валідний».
func NewIdentity(email, fullName string, now time.Time) (*Identity, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, ErrInvalidEmail
	}

	if fullName == "" {
		return nil, ErrEmptyName
	}

	return &Identity{
		Email:    email,
		FullName: fullName,
		Active:   true,
		Created:  now,
	}, nil
}
