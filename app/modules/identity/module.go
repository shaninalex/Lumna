package identity

import (
	"errors"
	"log/slog"

	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/modules/identity/contract"
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

	identities *storage.IdentityRepo
	// creds      *storage.CredentialRepo
	reader *infra.Reader

	// register     *usecase.Register
	// authenticate *usecase.Authenticate
	// getProfile   *usecase.GetProfile
}

func New(d Deps) *Module {
	identities := storage.NewIdentityRepo(d.DB)

	return &Module{
		deps:       d,
		identities: identities,
		// creds:      creds,
		reader: infra.NewReader(identities),

		// register:     usecase.NewRegister(identities, creds, hasher, d.Mailer, d.Clock),
		// authenticate: usecase.NewAuthenticate(identities, creds, hasher, tokens, d.Clock),
		// getProfile:   usecase.NewGetProfile(identities),
	}
}

func (m *Module) Name() string { return "identity" }

func (m *Module) Reader() contract.Reader {
	return m.reader
}

// Register — підписка на команди/запити/події. Помилки НЕ ковтаються.
func (m *Module) Register(a *core.App) error {
	return errors.Join(
	// bus.RegisterCommand(a.Commands, m.register.Handle),
	// bus.RegisterCommand(a.Commands, m.authenticate.Handle),
	// bus.RegisterQuery(a.Queries, m.getProfile.Handle),
	)
}
