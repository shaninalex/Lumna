package identity

import (
	"errors"
	"log/slog"

	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/bus"
	"gitlab.com/shaninalex/lumna/app/modules/identity/contract"
	"gitlab.com/shaninalex/lumna/app/modules/identity/handlers"
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
	reader      *infra.Reader

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
		reader:      infra.NewReader(identities),

		register:    handlers.NewRegister(identities, credentials, hasher, d.Mailer),
		profileList: handlers.NewListProfiles(identities),
	}
}

func (m *Module) Name() string { return "identity" }

func (m *Module) Reader() contract.Reader {
	return m.reader
}

// Register — subscribe on commands/queries/events. With error awareness
func (m *Module) Register(a *core.App) error {
	return errors.Join(
		bus.RegisterCommand(a.Commands, m.register.Handle),
		bus.RegisterQuery(a.Queries, m.profileList.Handle),
	)
}
