package securetoken

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type Hasher struct {
	secret []byte
}

func New(secret []byte) *Hasher {
	return &Hasher{
		secret: secret,
	}
}

func (s *Hasher) Sign(claims jwt.RegisteredClaims) (string, error) {
	signed, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(s.secret)
	if err != nil {
		return "", err
	}

	return signed, nil
}

func (s *Hasher) Token() (string, error) {
	raw := make([]byte, 32)
	if _, err := rand.Read(raw); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(raw), nil
}

func (s *Hasher) Unpack(token, issuer string) (*jwt.Token, error) {
	result, err := jwt.ParseWithClaims(
		token,
		&jwt.RegisteredClaims{},
		s.keyFunc,
		jwt.WithValidMethods([]string{"HS256"}),
		jwt.WithExpirationRequired(),
		jwt.WithIssuer(issuer),
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Hasher) HashRefresh(plain string) string {
	sum := sha256.Sum256([]byte(plain))
	return hex.EncodeToString(sum[:])
}

func (s *Hasher) keyFunc(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method")
	}
	return s.secret, nil
}
