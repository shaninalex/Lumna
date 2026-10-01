package storage

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
	"gitlab.com/shaninalex/lumna/app/platform/database"
	"gorm.io/gorm"
)

type ActivityRepo struct {
	db    *database.DB
	clock clock.Clock
}

var _ domain.ActivityRepo = (*ActivityRepo)(nil)

func NewActivityRepo(db *database.DB, clock clock.Clock) *ActivityRepo {
	return &ActivityRepo{
		db:    db,
		clock: clock,
	}
}

func (s *ActivityRepo) Create(ctx context.Context, activity domain.Activity) (domain.Activity, error) {
	record := activityToRecord(activity)
	record.ID = 0
	record.CreatedAt = s.clock.Now()
	if err := gorm.G[activityRecord](s.db.From(ctx)).Create(ctx, &record); err != nil {
		return domain.Activity{}, err
	}
	return activityToDomain(record), nil
}

func (s *ActivityRepo) ListById(ctx context.Context, entityId int, entityType string) ([]domain.Activity, error) {
	records, err := gorm.G[activityRecord](s.db.From(ctx)).
		Where("entity_id = ? and entity_type", entityId, entityType).
		Find(ctx)
	if err != nil {
		return nil, err
	}
	activitiesDomain := make([]domain.Activity, len(records))
	for i, record := range records {
		activitiesDomain[i] = activityToDomain(record)
	}
	return activitiesDomain, nil
}
