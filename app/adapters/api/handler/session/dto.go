package session

import (
	"time"

	"gitlab.com/shaninalex/lumna/app/modules/auth/contract"
)

type refreshTokenDTO struct {
	ID         int       `json:"id"`
	IdentityID int       `json:"identity_id"`
	ClientID   string    `json:"client_id"`
	Scopes     string    `json:"scopes"`
	ExpiresAt  time.Time `json:"expires_at"`
	Revoked    bool      `json:"revoked"`
	CreatedAt  time.Time `json:"created_at"`
}

func toRefreshTokenDTO(t contract.RefreshTokenView) refreshTokenDTO {
	return refreshTokenDTO{
		ID:         t.ID,
		IdentityID: t.IdentityID,
		ClientID:   t.ClientID,
		Scopes:     t.Scopes,
		ExpiresAt:  t.ExpiresAt,
		Revoked:    t.Revoked,
		CreatedAt:  t.CreatedAt,
	}
}
