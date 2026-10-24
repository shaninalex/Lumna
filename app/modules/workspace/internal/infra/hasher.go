package infra

import "gitlab.com/shaninalex/lumna/app/core/secure"

// TokenHasher implements domain.TokenHasher with opaque tokens: the plain
// token goes into the invitation link, only its hash is stored.
type TokenHasher struct{}

func NewTokenHasher() *TokenHasher {
	return &TokenHasher{}
}

func (s *TokenHasher) CreateToken() (string, string, error) {
	plain, err := secure.NewOpaque()
	if err != nil {
		return "", "", err
	}
	return plain, secure.HashOpaque(plain), nil
}

func (s *TokenHasher) Hash(token string) string {
	return secure.HashOpaque(token)
}
