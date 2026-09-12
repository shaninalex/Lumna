package workspace

import (
	"errors"
	"log/slog"

	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/bus"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/handlers"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/infra/storage"
	"gitlab.com/shaninalex/lumna/app/platform/database"
)

type Deps struct {
	DB  *database.DB
	Log *slog.Logger
}

type Module struct {
	deps Deps

	// repositories
	workspaceRepo         *storage.WorkspaceRepo
	identityWorkspaceRepo *storage.IdentityWorkspaceRepo

	// command handlers
	createWorkspace        *handlers.CreateWorkspace
	addIdentityToWorkspace *handlers.AddIdentityToWorkspace
}

func (m *Module) Name() string { return "workspace" }

func New(d Deps) *Module {
	workspaceRepo := storage.NewWorkspaceRepo(d.DB)
	identityWorkspaceRepo := storage.NewIdentityWorkspaceRepo(d.DB)
	return &Module{
		deps: d,

		workspaceRepo:          workspaceRepo,
		identityWorkspaceRepo:  identityWorkspaceRepo,
		createWorkspace:        handlers.NewCreateWorkspace(workspaceRepo),
		addIdentityToWorkspace: handlers.NewAddIdentityToWorkspace(identityWorkspaceRepo),
	}
}

func (m *Module) Register(a *core.App) error {
	return errors.Join(
		bus.RegisterCommand(a.Commands, m.createWorkspace.Handle),
		bus.RegisterCommand(a.Commands, m.addIdentityToWorkspace.Handle),
	)
}
