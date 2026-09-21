package storage

import (
	"database/sql"
	"time"

	"gitlab.com/shaninalex/lumna/app/modules/auth/internal/domain"
)

// refreshTokenRecord — shape in database
type refreshTokenRecord struct {
	ID         int            `gorm:"primaryKey;autoIncrement"`
	IdentityID int            `gorm:"column:identity_id;not null"`
	Hash       string         `gorm:"column:hash;not null;index"`
	ClientID   sql.NullString `gorm:"column:client_id;null"`
	Scopes     sql.NullString `gorm:"column:scopes;null"`
	ExpiresAt  time.Time      `gorm:"column:expires_at;not null"`
	Revoked    bool           `gorm:"column:revoked;not null;default:false"`
	CreatedAt  time.Time      `gorm:"column:created_at"`
}

func (refreshTokenRecord) TableName() string { return "refresh_tokens" }

func toDomainRefreshToken(rec refreshTokenRecord) domain.RefreshToken {
	return domain.RefreshToken{
		ID:         rec.ID,
		IdentityID: rec.IdentityID,
		Hash:       rec.Hash,
		ClientID:   rec.ClientID.String,
		Scopes:     rec.Scopes.String,
		ExpiresAt:  rec.ExpiresAt,
		Revoked:    rec.Revoked,
		CreatedAt:  rec.CreatedAt,
	}
}
