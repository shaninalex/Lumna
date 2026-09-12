package infra

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/core/actor"
	"gitlab.com/shaninalex/lumna/app/modules/auth/internal/domain"
)

// Verifier implements contract.Verifier.
type Verifier struct {
	tokens domain.Tokens
}

func NewVerifier(t domain.Tokens) *Verifier { return &Verifier{tokens: t} }

// Verify never reports why a token was rejected: expired, forged and
// malformed all look the same from outside, so a caller cannot probe.
func (v *Verifier) Verify(_ context.Context, accessToken string) (actor.Actor, error) {
	if accessToken == "" {
		return actor.Actor{}, domain.ErrBadAccessToken
	}

	identityID, err := v.tokens.ParseAccess(accessToken)
	if err != nil {
		return actor.Actor{}, domain.ErrBadAccessToken
	}

	return actor.Actor{IdentityID: identityID}, nil
}
