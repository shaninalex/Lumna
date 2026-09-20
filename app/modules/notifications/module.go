package notifications

import (
	"errors"
	"log/slog"

	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/bus"
	"gitlab.com/shaninalex/lumna/app/modules/notifications/internal/handlers"
	"gitlab.com/shaninalex/lumna/app/modules/notifications/internal/infra/storage"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
	"gitlab.com/shaninalex/lumna/app/platform/database"
)

type Deps struct {
	DB       *database.DB
	Log      *slog.Logger
	Clock    clock.Clock
	EventBus *bus.EventBus
}

type Module struct {
	deps Deps

	// event handlers
	workItemStageChanged *handlers.WorkItemStageChanged
}

func New(d Deps) *Module {
	repo := storage.NewNotificationRepo(d.DB)
	return &Module{
		deps: d,

		workItemStageChanged: handlers.NewWorkItemStageChanged(repo, d.Clock),
	}
}

func (s *Module) Name() string { return "notifications" }

func (s *Module) Register(a *core.App) error {
	return errors.Join(
		bus.Subscribe(a.Events, s.workItemStageChanged.Handle),
	)
}
