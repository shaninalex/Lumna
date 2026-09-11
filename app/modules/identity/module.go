package identity

import (
	"errors"
	"log/slog"

	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/bus"
	"gitlab.com/shaninalex/lumna/app/modules/identity/contract"
	"gitlab.com/shaninalex/lumna/app/modules/identity/internal/handlers"
	"gitlab.com/shaninalex/lumna/app/modules/identity/internal/infra"
	"gitlab.com/shaninalex/lumna/app/modules/identity/internal/infra/storage"
	"gitlab.com/shaninalex/lumna/app/platform/database"
)

type Deps struct {
	DB     *database.DB
	Log    *slog.Logger
	Secret []byte

	Mailer contract.Mailer // port, defined by other module
}

type Module struct {
	deps Deps

	identities  *storage.IdentityRepo
	credentials *storage.CredentialRepo

	// bridges — what this module exposes to other modules
	reader        *infra.Reader
	authenticator *infra.Authenticator
	provisioner   *infra.Provisioner

	// commands
	register *handlers.Register

	// query
	profileList *handlers.ListProfiles
}

func New(d Deps) *Module {
	identities := storage.NewIdentityRepo(d.DB)
	credentials := storage.NewCredentialRepo(d.DB)
	hasher := infra.NewHasher()

	return &Module{
		deps:        d,
		identities:  identities,
		credentials: credentials,

		reader:        infra.NewReader(identities),
		authenticator: infra.NewAuthenticator(identities, credentials, hasher),
		provisioner:   infra.NewProvisioner(identities),

		register:    handlers.NewRegister(identities, credentials, hasher, d.Mailer),
		profileList: handlers.NewListProfiles(identities),
	}
}

func (m *Module) Name() string { return "identity" }

// Reader — bridge. Read-only projections for other modules.
func (m *Module) Reader() contract.Reader {
	return m.reader
}

// Authenticator — bridge. Password check without exposing the hash.
func (m *Module) Authenticator() contract.Authenticator {
	return m.authenticator
}

// Provisioner — bridge. Create-if-missing, inside the caller's transaction.
func (m *Module) Provisioner() contract.Provisioner {
	return m.provisioner
}

// Register — subscribe on commands/queries/events. With error awareness
func (m *Module) Register(a *core.App) error {
	return errors.Join(
		bus.RegisterCommand(a.Commands, m.register.Handle),
		bus.RegisterQuery(a.Queries, m.profileList.Handle),
	)
}
