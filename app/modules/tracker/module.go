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
	stageRepo domain.StageRepo

	// commands
	createScope    *handlers.ScopeCreate
	createStage    *handlers.StateCreate
	createWorkItem *handlers.WorkItemCreate

	// queries
	scopeList *handlers.ScopeList
	stageList *handlers.StageList
}

func New(d Deps) *Module {
	scopeRepo := storage.NewScopeRepo(d.DB)
	stageRepo := storage.NewStageRepo(d.DB)

	return &Module{
		deps: d,

		scopeRepo:      scopeRepo,
		stageRepo:      stageRepo,
		createScope:    handlers.NewCreateScope(scopeRepo, d.Clock),
		createStage:    handlers.NewCreateStage(stageRepo, d.Clock),
		createWorkItem: handlers.NewWorkItemCreate(),
		scopeList:      handlers.NewScopeList(scopeRepo),
		stageList:      handlers.NewStageList(stageRepo),
	}
}

func (m *Module) Name() string { return "tracker" }

func (m *Module) Register(a *core.App) error {
	return errors.Join(
		bus.RegisterCommand(a.Commands, m.createScope.Handle),
		bus.RegisterCommand(a.Commands, m.createStage.Handle),
		bus.RegisterCommand(a.Commands, m.createWorkItem.Handle),
		bus.RegisterQuery(a.Queries, m.scopeList.Handle),
		bus.RegisterQuery(a.Queries, m.stageList.Handle),
	)
}
