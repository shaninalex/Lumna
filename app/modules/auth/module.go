package auth

import (
	"errors"
	"log/slog"
	"time"

	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/bus"
	"gitlab.com/shaninalex/lumna/app/modules/auth/contract"
	"gitlab.com/shaninalex/lumna/app/modules/auth/internal/handlers"
	"gitlab.com/shaninalex/lumna/app/modules/auth/internal/infra"
	"gitlab.com/shaninalex/lumna/app/modules/auth/internal/infra/storage"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
	"gitlab.com/shaninalex/lumna/app/platform/database"
)

type Deps struct {
	DB     *database.DB
	Log    *slog.Logger
	Clock  clock.Clock
	Secret []byte

	AccessTTL  time.Duration
	RefreshTTL time.Duration
	Issuer     string

	Identities  contract.IdentityAuthenticator
	Provisioner contract.IdentityProvisioner
}

type Module struct {
	deps Deps

	refresh *storage.RefreshTokenRepo

	// bridges — what this module exposes to adapters and other modules
	verifier *infra.Verifier

	// commands
	emailLogin     *handlers.EmailLogin
	refreshSession *handlers.RefreshSession
	logout         *handlers.Logout
}

func New(d Deps) *Module {
	refresh := storage.NewRefreshTokenRepo(d.DB)
	tokens := infra.NewTokens(d.Secret, d.AccessTTL, d.RefreshTTL, d.Issuer)

	return &Module{
		deps:    d,
		refresh: refresh,

		verifier: infra.NewVerifier(tokens),

		emailLogin:     handlers.NewEmailLogin(d.Identities, tokens, refresh, d.Clock),
		refreshSession: handlers.NewRefreshSession(tokens, refresh, d.Clock),
		logout:         handlers.NewLogout(tokens, refresh, d.Clock),
	}
}

func (m *Module) Name() string { return "auth" }

// Verifier — bridge. Access-token verification for the HTTP middleware.
func (m *Module) Verifier() contract.Verifier { return m.verifier }

// Register — subscribe on commands/queries/events. With error awareness
func (m *Module) Register(a *core.App) error {
	return errors.Join(
		bus.RegisterCommand(a.Commands, m.emailLogin.Handle),
		bus.RegisterCommand(a.Commands, m.refreshSession.Handle),
		bus.RegisterCommand(a.Commands, m.logout.Handle),
	)
}
