package workspace

import (
	"errors"
	"log/slog"

	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/bus"
	"gitlab.com/shaninalex/lumna/app/modules/identity/contract"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/handlers"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/infra"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/infra/storage"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
	"gitlab.com/shaninalex/lumna/app/platform/database"
	"gitlab.com/shaninalex/lumna/app/platform/email"
)

type Deps struct {
	DB     *database.DB
	Log    *slog.Logger
	Clock  clock.Clock
	Secret []byte

	Profiler    contract.Profiler
	EmailSender email.Sender
}

type Module struct {
	deps Deps

	// command handlers
	createWorkspace        *handlers.CreateWorkspace
	addIdentityToWorkspace *handlers.AddIdentityToWorkspace
	createProject          *handlers.CreateProject
	createInvitation       *handlers.CreateInvitation

	// query handlers
	workspaceList  *handlers.WorkspaceList
	projectList    *handlers.ProjectList
	membersList    *handlers.MembersList
	invitationList *handlers.InvitationList
}

func (m *Module) Name() string { return "workspace" }

func New(d Deps) *Module {
	workspaceRepo := storage.NewWorkspaceRepo(d.DB)
	identityWorkspaceRepo := storage.NewIdentityWorkspaceRepo(d.DB)
	projectRepo := storage.NewProjectRepo(d.DB)
	invitationRepo := storage.NewInvitationRepo(d.DB, d.Clock)
	hasher := infra.NewTokenHasher(d.Secret, "workspace")

	return &Module{
		deps: d,

		createWorkspace:        handlers.NewCreateWorkspace(workspaceRepo, d.Clock),
		addIdentityToWorkspace: handlers.NewAddIdentityToWorkspace(identityWorkspaceRepo, d.Clock),
		createProject:          handlers.NewCreateProject(projectRepo, d.Clock),
		createInvitation:       handlers.NewCreateInvitation(invitationRepo, hasher, d.Clock, d.EmailSender),

		workspaceList:  handlers.NewWorkspaceList(workspaceRepo),
		projectList:    handlers.NewProjectList(projectRepo),
		membersList:    handlers.NewMembersList(identityWorkspaceRepo, d.Profiler),
		invitationList: handlers.NewInvitationList(invitationRepo),
	}
}

func (m *Module) Register(a *core.App) error {
	return errors.Join(
		bus.RegisterCommand(a.Commands, m.createWorkspace.Handle),
		bus.RegisterCommand(a.Commands, m.addIdentityToWorkspace.Handle),
		bus.RegisterCommand(a.Commands, m.createProject.Handle),
		bus.RegisterCommand(a.Commands, m.createInvitation.Handle),
		bus.RegisterQuery(a.Queries, m.workspaceList.Handle),
		bus.RegisterQuery(a.Queries, m.projectList.Handle),
		bus.RegisterQuery(a.Queries, m.membersList.Handle),
		bus.RegisterQuery(a.Queries, m.invitationList.Handle),
	)
}
