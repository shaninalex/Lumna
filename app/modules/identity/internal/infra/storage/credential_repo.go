package storage

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/identity/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/database"
)

type CredentialRepo struct{ db *database.DB }

func NewCredentialRepo(db *database.DB) *CredentialRepo {
	return &CredentialRepo{db: db}
}

func (c2 CredentialRepo) Save(ctx context.Context, c *domain.Credential) error {
	//TODO implement me
	panic("implement me")
}

func (c2 CredentialRepo) ByIdentityAndProvider(ctx context.Context, id int, p domain.Provider) (*domain.Credential, error) {
	//TODO implement me
	panic("implement me")
}
