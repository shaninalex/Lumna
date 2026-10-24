package secure_test

import (
	"errors"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var testSecret = []byte("test-secret-key")

type payload struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var errBadPayload = errors.New("bad payload")

type validated struct {
	ID int `json:"id"`
}

func (v validated) Validate() error {
	if v.ID <= 0 {
		return errBadPayload
	}
	return nil
}

func newCodec[P any](kind string, ttl time.Duration) *secure.Codec[P] {
	return secure.NewCodec[P](secure.NewSigner(testSecret, "test-issuer"), kind, ttl)
}

func TestCodecRoundTrip(t *testing.T) {
	c := newCodec[payload]("thing", time.Hour)
	want := payload{ID: 42, Name: "a:b"}

	token, ttl, err := c.Issue(want, time.Now())
	if err != nil {
		t.Fatalf("Issue() error = %v", err)
	}
	if ttl != time.Hour {
		t.Errorf("ttl = %v, want %v", ttl, time.Hour)
	}

	got, err := c.Parse(token)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	if got != want {
		t.Errorf("Parse() = %+v, want %+v", got, want)
	}
}

func TestCodecRejectsOtherKind(t *testing.T) {
	signer := secure.NewSigner(testSecret, "test-issuer")
	access := secure.NewCodec[payload](signer, "access", time.Hour)
	invite := secure.NewCodec[payload](signer, "invitation", time.Hour)

	token, _, err := invite.Issue(payload{ID: 1}, time.Now())
	if err != nil {
		t.Fatalf("Issue() error = %v", err)
	}
	if _, err := access.Parse(token); err == nil {
		t.Fatal("Parse() accepted a token of another kind")
	}
}

func TestCodecRejectsExpired(t *testing.T) {
	c := newCodec[payload]("thing", time.Minute)

	token, _, err := c.Issue(payload{ID: 1}, time.Now().Add(-time.Hour))
	if err != nil {
		t.Fatalf("Issue() error = %v", err)
	}
	if _, err := c.Parse(token); err == nil {
		t.Fatal("Parse() accepted an expired token")
	}
}

func TestCodecRejectsWrongSecret(t *testing.T) {
	signer := newCodec[payload]("thing", time.Hour)
	verifier := secure.NewCodec[payload](secure.NewSigner([]byte("other"), "test-issuer"), "thing", time.Hour)

	token, _, err := signer.Issue(payload{ID: 1}, time.Now())
	if err != nil {
		t.Fatalf("Issue() error = %v", err)
	}
	if _, err := verifier.Parse(token); err == nil {
		t.Fatal("Parse() accepted a token signed with another secret")
	}
}

func TestCodecRejectsWrongIssuer(t *testing.T) {
	signer := newCodec[payload]("thing", time.Hour)
	verifier := secure.NewCodec[payload](secure.NewSigner(testSecret, "other-issuer"), "thing", time.Hour)

	token, _, err := signer.Issue(payload{ID: 1}, time.Now())
	if err != nil {
		t.Fatalf("Issue() error = %v", err)
	}
	if _, err := verifier.Parse(token); err == nil {
		t.Fatal("Parse() accepted a token from another issuer")
	}
}

func TestCodecRunsValidator(t *testing.T) {
	c := newCodec[validated]("thing", time.Hour)

	token, _, err := c.Issue(validated{ID: 0}, time.Now())
	if err != nil {
		t.Fatalf("Issue() error = %v", err)
	}
	if _, err := c.Parse(token); !errors.Is(err, errBadPayload) {
		t.Fatalf("Parse() error = %v, want %v", err, errBadPayload)
	}
}

func TestCodecRejectsMalformed(t *testing.T) {
	c := newCodec[payload]("thing", time.Hour)
	if _, err := c.Parse("not-a-jwt"); err == nil {
		t.Fatal("Parse() accepted a malformed token")
	}
}

// signRaw builds a token by hand, bypassing Codec, to forge the cases a
// Codec would never produce itself.
func signRaw(t *testing.T, method jwt.SigningMethod, key any, rc jwt.RegisteredClaims) string {
	t.Helper()
	token, err := jwt.NewWithClaims(method, rc).SignedString(key)
	if err != nil {
		t.Fatalf("SignedString() error = %v", err)
	}
	return token
}

func validRaw() jwt.RegisteredClaims {
	return jwt.RegisteredClaims{
		Issuer:    "test-issuer",
		Audience:  jwt.ClaimStrings{"thing"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
	}
}

func TestCodecRequiresExpiration(t *testing.T) {
	rc := validRaw()
	rc.ExpiresAt = nil
	token := signRaw(t, jwt.SigningMethodHS256, testSecret, rc)

	if _, err := newCodec[payload]("thing", time.Hour).Parse(token); err == nil {
		t.Fatal("Parse() accepted a token without exp")
	}
}

func TestCodecRejectsNonHS256(t *testing.T) {
	token := signRaw(t, jwt.SigningMethodHS384, testSecret, validRaw())

	if _, err := newCodec[payload]("thing", time.Hour).Parse(token); err == nil {
		t.Fatal("Parse() accepted HS384")
	}
}

func TestCodecRejectsNoneAlgorithm(t *testing.T) {
	token := signRaw(t, jwt.SigningMethodNone, jwt.UnsafeAllowNoneSignatureType, validRaw())

	if _, err := newCodec[payload]("thing", time.Hour).Parse(token); err == nil {
		t.Fatal("Parse() accepted alg=none")
	}
}
