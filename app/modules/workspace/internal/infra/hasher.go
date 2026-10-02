package infra

import (
	"gitlab.com/shaninalex/lumna/app/core/securetoken"
)

type TokenHasher struct {
	issuer string
	hasher *securetoken.Hasher
}

func NewTokenHasher(secret []byte, issuer string) *TokenHasher {
	return &TokenHasher{
		issuer: issuer,
		hasher: securetoken.New(secret),
	}
}

func (s *TokenHasher) CreateToken() (string, string, error) {
	plain, err := s.hasher.Token()
	if err != nil {
		return "", "", err
	}
	return plain, s.hasher.HashRefresh(plain), nil
}
