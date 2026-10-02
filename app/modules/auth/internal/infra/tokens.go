package infra

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gitlab.com/shaninalex/lumna/app/core/securetoken"
)

// Tokens implements domain.Tokens.
type Tokens struct {
	secret     []byte
	accessTTL  time.Duration
	refreshTTL time.Duration
	issuer     string

	hasher *securetoken.Hasher
}

func NewTokens(secret []byte, accessTTL, refreshTTL time.Duration, issuer string) *Tokens {
	return &Tokens{
		secret:     secret,
		accessTTL:  accessTTL,
		refreshTTL: refreshTTL,
		issuer:     issuer,

		hasher: securetoken.New(secret),
	}
}

func (t *Tokens) IssueAccess(identityID int, now time.Time) (string, time.Duration, error) {
	claims := jwt.RegisteredClaims{
		Subject:   strconv.Itoa(identityID),
		Issuer:    t.issuer,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(t.accessTTL)),
	}
	signed, err := t.hasher.Sign(claims)
	if err != nil {
		return "", 0, fmt.Errorf("auth: sign access token: %w", err)
	}
	return signed, t.accessTTL, nil
}

func (t *Tokens) IssueRefresh() (string, string, time.Duration, error) {
	plain, err := t.hasher.Token()
	if err != nil {
		return "", "", 0, fmt.Errorf("auth: generate refresh token: %w", err)
	}
	return plain, t.HashRefresh(plain), t.refreshTTL, nil
}

func (t *Tokens) HashRefresh(plain string) string {
	return t.hasher.HashRefresh(plain)
}

// ParseAccess validates the signature, issuer and expiry, and returns the
// identity the token was issued for.
func (t *Tokens) ParseAccess(accessToken string) (int, error) {
	token, err := t.hasher.Unpack(accessToken, t.issuer)
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return 0, fmt.Errorf("invalid token claims")
	}

	identityID, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return 0, fmt.Errorf("invalid subject")
	}

	return identityID, nil
}
