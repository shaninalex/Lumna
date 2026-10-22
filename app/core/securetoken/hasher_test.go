package securetoken_test

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gitlab.com/shaninalex/lumna/app/core/securetoken"
)

var testSecret = []byte("test-secret-key")

func TestSignAndUnpack(t *testing.T) {
	h := securetoken.New(testSecret)

	claims := jwt.RegisteredClaims{
		Issuer:    "test-issuer",
		Subject:   "user-123",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	signed, err := h.Sign(claims)
	if err != nil {
		t.Fatalf("Sign() error = %v", err)
	}
	if signed == "" {
		t.Fatal("Sign() returned an empty token")
	}

	parsed, err := h.Unpack(signed, "test-issuer")
	if err != nil {
		t.Fatalf("Unpack() error = %v", err)
	}
	if !parsed.Valid {
		t.Fatal("Unpack() returned an invalid token")
	}

	got, ok := parsed.Claims.(*jwt.RegisteredClaims)
	if !ok {
		t.Fatalf("claims type = %T, want *jwt.RegisteredClaims", parsed.Claims)
	}

	if got.Issuer != claims.Issuer {
		t.Errorf("Issuer = %q, want %q", got.Issuer, claims.Issuer)
	}
	if got.Subject != claims.Subject {
		t.Errorf("Subject = %q, want %q", got.Subject, claims.Subject)
	}
	if got.ExpiresAt == nil || !got.ExpiresAt.Equal(claims.ExpiresAt.Time) {
		t.Errorf("ExpiresAt = %v, want %v", got.ExpiresAt, claims.ExpiresAt)
	}
}

func TestUnpackRejectsWrongIssuer(t *testing.T) {
	h := securetoken.New(testSecret)

	token, err := h.Sign(jwt.RegisteredClaims{
		Issuer:    "expected-issuer",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
	})
	if err != nil {
		t.Fatalf("Sign() error = %v", err)
	}

	_, err = h.Unpack(token, "wrong-issuer")
	if err == nil {
		t.Fatal("Unpack() succeeded with the wrong issuer")
	}
}

func TestUnpackRejectsExpiredToken(t *testing.T) {
	h := securetoken.New(testSecret)

	token, err := h.Sign(jwt.RegisteredClaims{
		Issuer:    "test-issuer",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Hour)),
	})
	if err != nil {
		t.Fatalf("Sign() error = %v", err)
	}

	_, err = h.Unpack(token, "test-issuer")
	if err == nil {
		t.Fatal("Unpack() succeeded with an expired token")
	}
}

func TestUnpackRequiresExpiration(t *testing.T) {
	h := securetoken.New(testSecret)

	token, err := h.Sign(jwt.RegisteredClaims{
		Issuer: "test-issuer",
	})
	if err != nil {
		t.Fatalf("Sign() error = %v", err)
	}

	_, err = h.Unpack(token, "test-issuer")
	if err == nil {
		t.Fatal("Unpack() succeeded without an expiration claim")
	}
}

func TestUnpackRejectsWrongSecret(t *testing.T) {
	signer := securetoken.New(testSecret)
	verifier := securetoken.New([]byte("different-secret"))

	token, err := signer.Sign(jwt.RegisteredClaims{
		Issuer:    "test-issuer",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
	})
	if err != nil {
		t.Fatalf("Sign() error = %v", err)
	}

	_, err = verifier.Unpack(token, "test-issuer")
	if err == nil {
		t.Fatal("Unpack() succeeded with the wrong secret")
	}
}

func TestUnpackRejectsInvalidToken(t *testing.T) {
	h := securetoken.New(testSecret)

	_, err := h.Unpack("not-a-jwt", "test-issuer")
	if err == nil {
		t.Fatal("Unpack() succeeded with a malformed token")
	}
}

func TestUnpackRejectsNonHS256Algorithm(t *testing.T) {
	h := securetoken.New(testSecret)

	// HS384 is an HMAC method, but Unpack only allows HS256.
	token, err := jwt.NewWithClaims(
		jwt.SigningMethodHS384,
		jwt.RegisteredClaims{
			Issuer:    "test-issuer",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	).SignedString(testSecret)
	if err != nil {
		t.Fatalf("SignedString() error = %v", err)
	}

	_, err = h.Unpack(token, "test-issuer")
	if err == nil {
		t.Fatal("Unpack() accepted HS384; want HS256 only")
	}
}

func TestUnpackRejectsNoneAlgorithm(t *testing.T) {
	h := securetoken.New(testSecret)

	token := jwt.NewWithClaims(
		jwt.SigningMethodNone,
		jwt.RegisteredClaims{
			Issuer:    "test-issuer",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	)

	signed, err := token.SignedString(jwt.UnsafeAllowNoneSignatureType)
	if err != nil {
		t.Fatalf("SignedString() error = %v", err)
	}

	_, err = h.Unpack(signed, "test-issuer")
	if err == nil {
		t.Fatal("Unpack() accepted the none algorithm")
	}
}

func TestToken(t *testing.T) {
	h := securetoken.New(testSecret)

	first, err := h.Token()
	if err != nil {
		t.Fatalf("Token() error = %v", err)
	}
	second, err := h.Token()
	if err != nil {
		t.Fatalf("Token() second call error = %v", err)
	}

	if first == "" {
		t.Fatal("Token() returned an empty token")
	}
	if first == second {
		t.Fatal("Token() returned identical tokens on consecutive calls")
	}

	raw, err := base64.RawURLEncoding.DecodeString(first)
	if err != nil {
		t.Fatalf("Token() returned invalid raw URL-safe base64: %v", err)
	}
	if len(raw) != 32 {
		t.Errorf("decoded token length = %d, want 32 bytes", len(raw))
	}
	if strings.ContainsAny(first, "+/=") {
		t.Errorf("Token() contains non-URL-safe base64 characters: %q", first)
	}
}

func TestHashRefresh(t *testing.T) {
	h := securetoken.New(testSecret)

	plain := "refresh-token-example"
	got := h.HashRefresh(plain)

	sum := sha256.Sum256([]byte(plain))
	want := hex.EncodeToString(sum[:])

	if got != want {
		t.Errorf("HashRefresh() = %q, want %q", got, want)
	}
	if len(got) != 64 {
		t.Errorf("hash length = %d, want 64 hex characters", len(got))
	}
}

func TestHashRefreshDeterministic(t *testing.T) {
	h := securetoken.New(testSecret)

	plain := "same-refresh-token"

	first := h.HashRefresh(plain)
	second := h.HashRefresh(plain)

	if first != second {
		t.Errorf("HashRefresh() is not deterministic: %q != %q", first, second)
	}
}

func TestHashRefreshDifferentInputs(t *testing.T) {
	h := securetoken.New(testSecret)

	first := h.HashRefresh("token-one")
	second := h.HashRefresh("token-two")

	if first == second {
		t.Fatal("HashRefresh() returned the same hash for different inputs")
	}
}
