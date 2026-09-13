package tracker

import (
	"errors"
	"log/slog"

	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/bus"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/handlers"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/infra/storage"
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

	// repositories
	scopeRepo domain.ScopeRepo

	// commands
	createScope *handlers.CreateScope
}

func New(d Deps) *Module {
	scopeRepo := storage.NewScopeRepo(d.DB)

	return &Module{
		deps: d,

		scopeRepo:   scopeRepo,
		createScope: handlers.NewCreateScope(scopeRepo),
	}
}

func (m *Module) Name() string { return "tracker" }

func (m *Module) Register(a *core.App) error {
	return errors.Join(
		bus.RegisterCommand(a.Commands, m.createScope.Handle),
	)
}
