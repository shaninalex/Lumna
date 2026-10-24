package secure

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Validator interface {
	Validate() error
}

type Codec[P any] struct {
	signer *Signer
	kind   string
	ttl    time.Duration
}

func NewCodec[P any](s *Signer, kind string, ttl time.Duration) *Codec[P] {
	return &Codec[P]{
		signer: s,
		kind:   kind,
		ttl:    ttl,
	}
}

func (c *Codec[P]) Issue(p P, now time.Time) (token string, ttl time.Duration, err error) {
	cl := claims[P]{
		//jwt.RegisteredClaims
		Issuer:    c.signer.issuer,
		Audience:  jwt.ClaimStrings{c.kind},
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(c.ttl)),

		//payload
		Payload: p,
	}
	signed, err := c.signer.sign(cl)
	if err != nil {
		return "", 0, fmt.Errorf("securetoken: sign %s token: %w", c.kind, err)
	}
	return signed, c.ttl, nil
}

func (c *Codec[P]) Parse(token string) (P, error) {
	var zero P

	var cl claims[P]
	if err := c.signer.parse(token, &cl, c.kind); err != nil {
		return zero, fmt.Errorf("securetoken: parse %s token: %w", c.kind, err)
	}

	if v, ok := any(cl.Payload).(Validator); ok {
		if err := v.Validate(); err != nil {
			return zero, fmt.Errorf("securetoken: %s payload: %w", c.kind, err)
		}
	}

	return cl.Payload, nil
}

type claims[P any] struct {
	jwt.RegisteredClaims
	Payload P `json:"p"`
}
