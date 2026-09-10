package storage

import (
	"context"
	"errors"
	"strings"
	"time"

	"gitlab.com/shaninalex/lumna/app/modules/identity/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/database"
	"gorm.io/gorm"
)

type IdentityRepo struct {
	db *database.DB
}

func NewIdentityRepo(db *database.DB) *IdentityRepo { return &IdentityRepo{db: db} }

var _ domain.IdentityRepo = (*IdentityRepo)(nil)

// ByEmail implements [domain.IdentityRepo].
func (r *IdentityRepo) ByEmail(ctx context.Context, email string) (*domain.Identity, error) {
	rec, err := gorm.G[identityRecord](r.db.From(ctx)).Where("email = ?", strings.ToLower(email)).First(ctx)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return toDomainIdentity(rec), nil
}

// DisplayNames implements [domain.IdentityRepo].
func (r *IdentityRepo) DisplayNames(ctx context.Context, ids []int) (map[int]string, error) {
	if len(ids) == 0 {
		return map[int]string{}, nil
	}
	recs, err := gorm.G[identityRecord](r.db.From(ctx)).Where("id IN ?", ids).Find(ctx)
	if err != nil {
		return nil, err
	}
	out := make(map[int]string, len(recs))
	for _, rec := range recs {
		out[rec.ID] = rec.FullName
	}
	return out, nil
}

// ByID implements [domain.IdentityRepo].
func (r *IdentityRepo) ByID(ctx context.Context, id int) (*domain.Identity, error) {
	rec, err := gorm.G[identityRecord](r.db.From(ctx)).Where("id = ?", id).First(ctx)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return toDomainIdentity(rec), nil
}

// Save implements [domain.IdentityRepo].
func (r *IdentityRepo) Save(ctx context.Context, i *domain.Identity) error {
	_, err := r.db.InTx(ctx, func(ctx context.Context) (any, error) {
		record := identityRecord{
			ID:        i.ID,
			Email:     i.Email,
			FullName:  i.FullName,
			Active:    i.Active,
			CreatedAt: time.Now(),
		}
		if err := r.db.From(ctx).Create(&record).Error; err != nil {
			return nil, err
		}

		return toDomainIdentity(record), nil
	})

	return err
}

func (r *IdentityRepo) List(ctx context.Context, limit, offset int) ([]*domain.Identity, error) {
	rec, err := gorm.G[identityRecord](r.db.From(ctx)).Limit(limit).Offset(offset).Find(ctx)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	result := []*domain.Identity{}
	for _, rec := range rec {
		result = append(result, toDomainIdentity(rec))
	}

	return result, nil
}
