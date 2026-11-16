package journal

import (
	"log/slog"

	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/bus"
	"gitlab.com/shaninalex/lumna/app/modules/journal/internal/handlers"
	"gitlab.com/shaninalex/lumna/app/modules/journal/internal/infra/storage"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
	"gitlab.com/shaninalex/lumna/app/platform/database"
)

type Deps struct {
	DB    *database.DB
	Log   *slog.Logger
	Clock clock.Clock
}

type Module struct {
	deps Deps

	// event handlers
	record *handlers.Record
}

func New(d Deps) *Module {
	entries := storage.NewEntryRepo(d.DB)
	return &Module{
		deps:   d,
		record: handlers.NewRecord(entries, d.Clock),
	}
}

func (m *Module) Name() string { return "journal" }

func (m *Module) Register(a *core.App) error {
	return bus.SubscribeTo(a.Events, m.record.Handle)
}
