package infra

import (
	"gitlab.com/shaninalex/lumna/app/core/securetoken"
)

// TokenHasher implements domain.TokenHasher with opaque tokens: the plain
// token goes into the invitation link, only its hash is stored.
type TokenHasher struct{}

func NewTokenHasher() *TokenHasher {
	return &TokenHasher{}
}

func (s *TokenHasher) CreateToken() (string, string, error) {
	plain, err := securetoken.NewOpaque()
	if err != nil {
		return "", "", err
	}
	return plain, securetoken.HashOpaque(plain), nil
}

func (s *TokenHasher) Hash(token string) string {
	return securetoken.HashOpaque(token)
}
