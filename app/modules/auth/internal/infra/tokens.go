package infra

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Tokens implements domain.Tokens.
type Tokens struct {
	secret     []byte
	accessTTL  time.Duration
	refreshTTL time.Duration
	issuer     string
}

func NewTokens(secret []byte, accessTTL, refreshTTL time.Duration, issuer string) *Tokens {
	return &Tokens{
		secret:     secret,
		accessTTL:  accessTTL,
		refreshTTL: refreshTTL,
		issuer:     issuer,
	}
}

func (t *Tokens) IssueAccess(identityID int, now time.Time) (string, time.Duration, error) {
	claims := jwt.RegisteredClaims{
		Subject:   strconv.Itoa(identityID),
		Issuer:    t.issuer,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(t.accessTTL)),
	}

	signed, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(t.secret)
	if err != nil {
		return "", 0, fmt.Errorf("auth: sign access token: %w", err)
	}

	return signed, t.accessTTL, nil
}

func (t *Tokens) IssueRefresh() (string, string, time.Duration, error) {
	raw := make([]byte, 32)
	if _, err := rand.Read(raw); err != nil {
		return "", "", 0, fmt.Errorf("auth: generate refresh token: %w", err)
	}

	plain := base64.RawURLEncoding.EncodeToString(raw)
	return plain, t.HashRefresh(plain), t.refreshTTL, nil
}

func (t *Tokens) HashRefresh(plain string) string {
	sum := sha256.Sum256([]byte(plain))
	return hex.EncodeToString(sum[:])
}

func (t *Tokens) keyFunc(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method")
	}
	return t.secret, nil
}

// ParseAccess validates the signature, issuer and expiry, and returns the
// identity the token was issued for.
func (t *Tokens) ParseAccess(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&jwt.RegisteredClaims{},
		t.keyFunc,
		jwt.WithValidMethods([]string{"HS256"}),
		jwt.WithExpirationRequired(),
		jwt.WithIssuer(t.issuer),
	)
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
