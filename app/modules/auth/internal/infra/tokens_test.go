package infra_test

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"
	"time"

	"gitlab.com/shaninalex/lumna/app/modules/auth/internal/infra"
)

func newTestTokens() *infra.Tokens {
	return infra.NewTokens(
		[]byte("test-secret-key"),
		15*time.Minute,
		24*time.Hour,
		"test-issuer",
	)
}

func TestParseAccess(t *testing.T) {
	tokens := newTestTokens()
	now := time.Now().UTC()
	identityID := 456

	signed, _, err := tokens.IssueAccess(identityID, now)
	if err != nil {
		t.Fatalf("IssueAccess() error = %v", err)
	}

	got, err := tokens.ParseAccess(signed)
	if err != nil {
		t.Fatalf("ParseAccess() error = %v", err)
	}
	if got != identityID {
		t.Errorf("ParseAccess() = %d, want %d", got, identityID)
	}
}

func TestParseAccessRejectsInvalidToken(t *testing.T) {
	tokens := newTestTokens()

	_, err := tokens.ParseAccess("not-a-jwt")
	if err == nil {
		t.Fatal("ParseAccess() succeeded with an invalid token")
	}
}

func TestParseAccessRejectsWrongSecret(t *testing.T) {
	signer := newTestTokens()
	verifier := infra.NewTokens(
		[]byte("different-secret"),
		15*time.Minute,
		24*time.Hour,
		"test-issuer",
	)

	signed, _, err := signer.IssueAccess(123, time.Now().UTC())
	if err != nil {
		t.Fatalf("IssueAccess() error = %v", err)
	}

	_, err = verifier.ParseAccess(signed)
	if err == nil {
		t.Fatal("ParseAccess() succeeded with the wrong secret")
	}
}

func TestParseAccessRejectsWrongIssuer(t *testing.T) {
	signer := newTestTokens()
	verifier := infra.NewTokens(
		[]byte("test-secret-key"),
		15*time.Minute,
		24*time.Hour,
		"another-issuer",
	)

	signed, _, err := signer.IssueAccess(123, time.Now().UTC())
	if err != nil {
		t.Fatalf("IssueAccess() error = %v", err)
	}

	_, err = verifier.ParseAccess(signed)
	if err == nil {
		t.Fatal("ParseAccess() succeeded with the wrong issuer")
	}
}

func TestParseAccessRejectsExpiredToken(t *testing.T) {
	tokens := newTestTokens()
	now := time.Now().UTC().Add(-time.Hour)

	signed, _, err := tokens.IssueAccess(123, now)
	if err != nil {
		t.Fatalf("IssueAccess() error = %v", err)
	}

	_, err = tokens.ParseAccess(signed)
	if err == nil {
		t.Fatal("ParseAccess() succeeded with an expired token")
	}
}

func TestIssueRefresh(t *testing.T) {
	tokens := newTestTokens()

	plain, hash, _, err := tokens.IssueRefresh()
	if err != nil {
		t.Fatalf("IssueRefresh() error = %v", err)
	}

	if plain == "" {
		t.Fatal("IssueRefresh() returned an empty token")
	}
	if hash == "" {
		t.Fatal("IssueRefresh() returned an empty hash")
	}

	sum := sha256.Sum256([]byte(plain))
	wantHash := hex.EncodeToString(sum[:])
	if hash != wantHash {
		t.Errorf("hash = %q, want %q", hash, wantHash)
	}

	if len(hash) != 64 {
		t.Errorf("hash length = %d, want 64", len(hash))
	}

	plain2, hash2, _, err := tokens.IssueRefresh()
	if err != nil {
		t.Fatalf("IssueRefresh() second call error = %v", err)
	}
	if plain == plain2 {
		t.Fatal("IssueRefresh() generated duplicate refresh tokens")
	}
	if hash == hash2 {
		t.Fatal("IssueRefresh() generated duplicate refresh hashes")
	}
}

func TestHashRefresh(t *testing.T) {
	tokens := newTestTokens()
	plain := "refresh-token-example"

	got := tokens.HashRefresh(plain)

	sum := sha256.Sum256([]byte(plain))
	want := hex.EncodeToString(sum[:])

	if got != want {
		t.Errorf("HashRefresh() = %q, want %q", got, want)
	}
}

func TestHashRefreshDeterministic(t *testing.T) {
	tokens := newTestTokens()
	plain := "same-refresh-token"

	first := tokens.HashRefresh(plain)
	second := tokens.HashRefresh(plain)

	if first != second {
		t.Errorf("HashRefresh() is not deterministic: %q != %q", first, second)
	}
}

func TestHashRefreshDifferentInputs(t *testing.T) {
	tokens := newTestTokens()

	first := tokens.HashRefresh("token-one")
	second := tokens.HashRefresh("token-two")

	if first == second {
		t.Fatal("HashRefresh() returned identical hashes for different inputs")
	}
}
