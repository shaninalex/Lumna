package infra

import (
	"golang.org/x/crypto/bcrypt"
)

type Hasher struct {
}

func (s *Hasher) Hash(plain string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plain), 12)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (s *Hasher) Verify(hash, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
}
