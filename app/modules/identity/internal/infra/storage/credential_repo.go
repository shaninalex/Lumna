package storage

import (
	"context"
	"errors"

	"gitlab.com/shaninalex/lumna/app/modules/identity/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/database"
	"gorm.io/gorm"
)

type CredentialRepo struct{ db *database.DB }

func NewCredentialRepo(db *database.DB) *CredentialRepo {
	return &CredentialRepo{db: db}
}

var _ domain.CredentialRepo = (*CredentialRepo)(nil)

// Save implements [domain.CredentialRepo].
//
// No transaction here: the write chain owns it, From(ctx) hands it over.
func (r *CredentialRepo) Save(ctx context.Context, c *domain.Credential) error {
	record := credentialRecord{
		ID:             c.ID,
		IdentityID:     c.IdentityID,
		Provider:       c.Provider,
		ProviderUserID: c.ProviderUserID,
		Email:          c.Email,
		PasswordHash:   c.PasswordHash,
		CreatedAt:      c.CreatedAt,
	}
	if err := r.db.From(ctx).Save(&record).Error; err != nil {
		return err
	}

	c.ID = record.ID
	return nil
}

// ByIdentityAndProvider implements [domain.CredentialRepo].
func (r *CredentialRepo) ByIdentityAndProvider(ctx context.Context, id int, p domain.Provider) (*domain.Credential, error) {
	rec, err := gorm.G[credentialRecord](r.db.From(ctx)).
		Where("identity_id = ? and provider = ?", id, string(p)).
		First(ctx)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return toDomainCredential(rec), nil
}
