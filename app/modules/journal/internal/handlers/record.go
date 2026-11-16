package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	"gitlab.com/shaninalex/lumna/app/modules/journal/contract"
	"gitlab.com/shaninalex/lumna/app/modules/journal/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

// Record stores every [contract.Recordable] event as a journal entry.
type Record struct {
	repo  domain.EntryRepo
	clock clock.Clock
}

func NewRecord(repo domain.EntryRepo, clock clock.Clock) *Record {
	return &Record{repo: repo, clock: clock}
}

func (s *Record) Handle(ctx context.Context, e contract.Recordable) error {
	data, err := json.Marshal(e)
	if err != nil {
		return fmt.Errorf("journal: marshal %s: %w", e.EventName(), err)
	}
	entityType, entityID := e.JournalSubject()
	_, err = s.repo.Save(ctx, domain.Entry{
		ActorID:    e.JournalActor(),
		EntityType: entityType,
		EntityID:   entityID,
		EventType:  e.EventName(),
		Data:       string(data),
		CreatedAt:  s.clock.Now(),
	})
	return err
}
