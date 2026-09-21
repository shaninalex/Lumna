package storage

import (
	"database/sql"
	"time"

	"gitlab.com/shaninalex/lumna/app/modules/identity/internal/domain"
)

// identityRecord — shape in database
type identityRecord struct {
	ID        int          `gorm:"primaryKey;autoIncrement"` //nolint:gofmt
	Email     string       `gorm:"uniqueIndex;not null"`
	FullName  string       `gorm:"column:full_name;not null"`
	Active    bool         `gorm:"default:true"`
	CreatedAt time.Time    `gorm:"autoCreateTime"`
	UpdatedAt sql.NullTime `gorm:"column:updated_at"`
}

func (identityRecord) TableName() string { return "identities" }

func toDomainIdentity(rec identityRecord) domain.Identity {
	return domain.Identity{
		ID:       rec.ID,
		Email:    rec.Email,
		FullName: rec.FullName,
		Active:   rec.Active,
		Created:  rec.CreatedAt,
		Updated:  rec.UpdatedAt.Time,
	}
}

// credentialRecord — shape in database
type credentialRecord struct {
	ID             int            `gorm:"primaryKey;autoIncrement"`
	IdentityID     int            `gorm:"column:identity_id;not null"`
	Provider       string         `gorm:"column:provider;not null"`
	ProviderUserID sql.NullString `gorm:"column:provider_user_id;null"`
	Email          sql.NullString `gorm:"column:email;null"`
	PasswordHash   sql.NullString `gorm:"column:password_hash;null"`
	CreatedAt      time.Time      `gorm:"autoCreateTime"`
}

func (credentialRecord) TableName() string { return "credentials" }

func toDomainCredential(rec credentialRecord) domain.Credential {
	return domain.Credential{
		ID:             rec.ID,
		IdentityID:     rec.IdentityID,
		Provider:       rec.Provider,
		ProviderUserID: rec.ProviderUserID.String,
		Email:          rec.Email.String,
		PasswordHash:   rec.PasswordHash.String,
		CreatedAt:      rec.CreatedAt,
	}
}
