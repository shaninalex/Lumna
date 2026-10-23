package infra_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"gitlab.com/shaninalex/lumna/app/core/actor"
	"gitlab.com/shaninalex/lumna/app/core/securetoken"
	"gitlab.com/shaninalex/lumna/app/modules/auth/contract"
	"gitlab.com/shaninalex/lumna/app/modules/auth/internal/domain"
	"gitlab.com/shaninalex/lumna/app/modules/auth/internal/infra"
)

type fixture struct {
	access     *securetoken.Codec[domain.AccessClaims]
	invitation *securetoken.Codec[contract.InvitationClaims]

	accessVerifier     contract.Verifier[actor.Actor]
	invitationVerifier contract.Verifier[contract.InvitationClaims]
}

func newFixture() fixture {
	signer := securetoken.NewSigner([]byte("test-secret-key"), "test-issuer")
	f := fixture{
		access:     securetoken.NewCodec[domain.AccessClaims](signer, "access", 15*time.Minute),
		invitation: securetoken.NewCodec[contract.InvitationClaims](signer, "invitation", 15*time.Minute),
	}
	f.accessVerifier = infra.NewVerifier(f.access, func(c domain.AccessClaims) actor.Actor {
		return actor.Actor{IdentityID: c.IdentityID}
	})
	f.invitationVerifier = infra.NewVerifier(f.invitation, infra.Identity[contract.InvitationClaims])
	return f
}

func TestAccessVerifier(t *testing.T) {
	f := newFixture()

	token, _, err := f.access.Issue(domain.AccessClaims{IdentityID: 456}, time.Now())
	if err != nil {
		t.Fatalf("Issue() error = %v", err)
	}

	got, err := f.accessVerifier.Verify(context.Background(), token)
	if err != nil {
		t.Fatalf("Verify() error = %v", err)
	}
	if got.IdentityID != 456 {
		t.Errorf("IdentityID = %d, want 456", got.IdentityID)
	}
}

func TestInvitationVerifier(t *testing.T) {
	f := newFixture()
	want := contract.InvitationClaims{Email: "new@example.com", Role: "member"}

	token, _, err := f.invitation.Issue(want, time.Now())
	if err != nil {
		t.Fatalf("Issue() error = %v", err)
	}

	got, err := f.invitationVerifier.Verify(context.Background(), token)
	if err != nil {
		t.Fatalf("Verify() error = %v", err)
	}
	if got != want {
		t.Errorf("Verify() = %+v, want %+v", got, want)
	}
}

func TestVerifiersRejectEachOthersTokens(t *testing.T) {
	f := newFixture()
	ctx := context.Background()

	invitation, _, err := f.invitation.Issue(contract.InvitationClaims{Email: "a@b.c", Role: "admin"}, time.Now())
	if err != nil {
		t.Fatalf("Issue() error = %v", err)
	}
	if _, err := f.accessVerifier.Verify(ctx, invitation); !errors.Is(err, domain.ErrBadAccessToken) {
		t.Errorf("access verifier on invitation token: err = %v, want %v", err, domain.ErrBadAccessToken)
	}

	access, _, err := f.access.Issue(domain.AccessClaims{IdentityID: 1}, time.Now())
	if err != nil {
		t.Fatalf("Issue() error = %v", err)
	}
	if _, err := f.invitationVerifier.Verify(ctx, access); !errors.Is(err, domain.ErrBadAccessToken) {
		t.Errorf("invitation verifier on access token: err = %v, want %v", err, domain.ErrBadAccessToken)
	}
}

// Every rejection is the same error, so a caller cannot tell why.
func TestVerifierRejections(t *testing.T) {
	f := newFixture()

	expired, _, err := f.access.Issue(domain.AccessClaims{IdentityID: 1}, time.Now().Add(-time.Hour))
	if err != nil {
		t.Fatalf("Issue() error = %v", err)
	}
	zeroID, _, err := f.access.Issue(domain.AccessClaims{IdentityID: 0}, time.Now())
	if err != nil {
		t.Fatalf("Issue() error = %v", err)
	}

	for name, token := range map[string]string{
		"empty":     "",
		"malformed": "not-a-jwt",
		"expired":   expired,
		"zero id":   zeroID,
	} {
		t.Run(name, func(t *testing.T) {
			if _, err := f.accessVerifier.Verify(context.Background(), token); !errors.Is(err, domain.ErrBadAccessToken) {
				t.Errorf("Verify() err = %v, want %v", err, domain.ErrBadAccessToken)
			}
		})
	}
}

func TestInvitationClaimsValidate(t *testing.T) {
	for name, tc := range map[string]struct {
		claims  contract.InvitationClaims
		wantErr bool
	}{
		"valid":            {contract.InvitationClaims{Email: "a@b.c", Role: "member"}, false},
		"quoted local":     {contract.InvitationClaims{Email: `"a:b"@b.c`, Role: "member"}, false},
		"any role is fine": {contract.InvitationClaims{Email: "a@b.c", Role: "whatever"}, false},
		"bad email":        {contract.InvitationClaims{Email: "nope", Role: "member"}, true},
		"empty role":       {contract.InvitationClaims{Email: "a@b.c"}, true},
	} {
		t.Run(name, func(t *testing.T) {
			if err := tc.claims.Validate(); (err != nil) != tc.wantErr {
				t.Errorf("Validate() err = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
