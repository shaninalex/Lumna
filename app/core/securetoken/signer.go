package securetoken

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type Signer struct {
	secret []byte
	issuer string
}

func NewSigner(secret []byte, issuer string) *Signer {
	return &Signer{
		secret: secret,
		issuer: issuer,
	}
}

func (s *Signer) sign(c jwt.Claims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(s.secret)
}

func (s *Signer) parse(token string, dst jwt.Claims, audience string) error {
	_, err := jwt.ParseWithClaims(
		token,
		dst,
		s.keyFunc,
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithIssuer(s.issuer),
		jwt.WithAudience(audience),
		jwt.WithExpirationRequired(),
	)
	return err
}

func (s *Signer) keyFunc(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("securetoken: unexpected signing method %v", token.Header["alg"])
	}
	return s.secret, nil
}
