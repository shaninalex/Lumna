package secure_test

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"testing"

	"gitlab.com/shaninalex/lumna/app/core/secure"
)

func TestNewOpaque(t *testing.T) {
	first, err := secure.NewOpaque()
	if err != nil {
		t.Fatalf("NewOpaque() error = %v", err)
	}
	second, err := secure.NewOpaque()
	if err != nil {
		t.Fatalf("NewOpaque() error = %v", err)
	}
	if first == second {
		t.Fatal("NewOpaque() returned the same token twice")
	}

	raw, err := base64.RawURLEncoding.DecodeString(first)
	if err != nil {
		t.Fatalf("token is not raw URL-safe base64: %v", err)
	}
	if len(raw) != 32 {
		t.Errorf("decoded length = %d, want 32", len(raw))
	}
}

func TestHashOpaque(t *testing.T) {
	sum := sha256.Sum256([]byte("plain"))
	want := hex.EncodeToString(sum[:])

	if got := secure.HashOpaque("plain"); got != want {
		t.Errorf("HashOpaque() = %q, want %q", got, want)
	}
	if secure.HashOpaque("a") == secure.HashOpaque("b") {
		t.Error("HashOpaque() collides on different inputs")
	}
}
