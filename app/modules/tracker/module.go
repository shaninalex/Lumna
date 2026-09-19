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
	scopeRepo    domain.ScopeRepo
	stageRepo    domain.StageRepo
	workItemRepo domain.WorkingItemRepo

	// commands
	createScope        *handlers.ScopeCreate
	createStage        *handlers.StateCreate
	stageMove          *handlers.StageMove
	stageDelete        *handlers.StageDelete
	createWorkItem     *handlers.WorkItemCreate
	workItemMove       *handlers.WorkItemMove
	workItemTransfer   *handlers.WorkItemTransfer
	workItemUpdate     *handlers.WorkItemUpdate
	workItemAssignment *handlers.WorkItemAssignment
	workItemDelete     *handlers.WorkItemDelete

	// queries
	scopeList    *handlers.ScopeList
	stageList    *handlers.StageList
	workItemList *handlers.WorkItemList
}

func New(d Deps) *Module {
	scopeRepo := storage.NewScopeRepo(d.DB)
	stageRepo := storage.NewStageRepo(d.DB)
	workItemRepo := storage.NewWorkingItemRepo(d.DB)

	return &Module{
		deps: d,

		scopeRepo:          scopeRepo,
		stageRepo:          stageRepo,
		workItemRepo:       workItemRepo,
		createScope:        handlers.NewCreateScope(scopeRepo, d.Clock),
		createStage:        handlers.NewCreateStage(stageRepo, d.Clock),
		stageDelete:        handlers.NewStageDelete(workItemRepo, stageRepo),
		scopeList:          handlers.NewScopeList(scopeRepo),
		stageList:          handlers.NewStageList(stageRepo),
		stageMove:          handlers.NewStageMove(stageRepo, d.Clock),
		createWorkItem:     handlers.NewWorkItemCreate(workItemRepo, d.Clock),
		workItemList:       handlers.NewWorkItemList(workItemRepo, d.Clock),
		workItemMove:       handlers.NewWorkItemMove(workItemRepo, d.Clock),
		workItemTransfer:   handlers.NewWorkItemTransfer(workItemRepo, stageRepo, d.Clock),
		workItemUpdate:     handlers.NewWorkItemUpdate(workItemRepo, d.Clock),
		workItemAssignment: handlers.NewWorkItemAssignment(workItemRepo, d.Clock),
		workItemDelete:     handlers.NewWorkItemDelete(workItemRepo),
	}
}

func (m *Module) Name() string { return "tracker" }

func (m *Module) Register(a *core.App) error {
	return errors.Join(
		bus.RegisterCommand(a.Commands, m.createScope.Handle),
		bus.RegisterQuery(a.Queries, m.scopeList.Handle),

		bus.RegisterCommand(a.Commands, m.createStage.Handle),
		bus.RegisterCommand(a.Commands, m.stageMove.Handle),
		bus.RegisterCommand(a.Commands, m.stageDelete.Handle),
		bus.RegisterQuery(a.Queries, m.stageList.Handle),

		bus.RegisterCommand(a.Commands, m.createWorkItem.Handle),
		bus.RegisterCommand(a.Commands, m.workItemMove.Handle),
		bus.RegisterCommand(a.Commands, m.workItemTransfer.Handle),
		bus.RegisterCommand(a.Commands, m.workItemUpdate.Handle),
		bus.RegisterCommand(a.Commands, m.workItemAssignment.Handle),
		bus.RegisterCommand(a.Commands, m.workItemDelete.Handle),
		bus.RegisterQuery(a.Queries, m.workItemList.Handle),
	)
}
