package auth

import (
	"errors"
	"log/slog"
	"time"

	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/actor"
	"gitlab.com/shaninalex/lumna/app/core/bus"
	"gitlab.com/shaninalex/lumna/app/core/secure"
	"gitlab.com/shaninalex/lumna/app/modules/auth/contract"
	"gitlab.com/shaninalex/lumna/app/modules/auth/internal/domain"
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

	// InvitationTTL is how long an invitation session lasts: the time a user
	// has to set a password after opening the invitation link.
	InvitationTTL time.Duration

	Identities  contract.IdentityAuthenticator
	Provisioner contract.IdentityProvisioner
}

type Module struct {
	deps Deps

	// bridges — what this module exposes to adapters and other modules
	accessVerifier     contract.Verifier[actor.Actor]
	invitationVerifier contract.Verifier[contract.InvitationClaims]
	passwordVerifier   contract.PasswordVerifier

	// commands
	emailLogin        *handlers.EmailLogin
	refreshSession    *handlers.RefreshSession
	logout            *handlers.Logout
	invitationSession *handlers.InvitationSession

	// Workers
	refreshTokensCleaner *infra.RefreshTokensCleaner
}

// Token kinds. Each one is the `aud` claim of its tokens, so a token of one
// kind is rejected by the codec of any other. Never reuse or rename a kind
// casually: renaming one invalidates every token of that kind in flight.
//
// A new kind = a payload type (with Validate), a constant here, a codec and,
// if an adapter needs to check it, a verifier bridge below.
const (
	kindAccess     = "access"
	kindInvitation = "invitation"
)

func New(d Deps) *Module {
	refreshRepo := storage.NewRefreshTokenRepo(d.DB)

	// One signer for every signed kind; the codecs keep the kinds apart.
	signer := secure.NewSigner(d.Secret, d.Issuer)
	access := secure.NewCodec[domain.AccessClaims](signer, kindAccess, d.AccessTTL)
	invitation := secure.NewCodec[contract.InvitationClaims](signer, kindInvitation, d.InvitationTTL)

	refresh := infra.NewRefreshTokens(d.RefreshTTL, refreshRepo, access)
	sessions := handlers.NewSessionIssuer(access, refresh, refreshRepo)

	return &Module{
		deps: d,

		// bridges
		accessVerifier: infra.NewVerifier(access, func(c domain.AccessClaims) actor.Actor {
			return actor.Actor{IdentityID: c.IdentityID}
		}),
		invitationVerifier: infra.NewVerifier(invitation, infra.Identity[contract.InvitationClaims]),
		passwordVerifier:   infra.NewPasswordVerifier(),

		// commands
		emailLogin:        handlers.NewEmailLogin(d.Identities, sessions, d.Clock),
		refreshSession:    handlers.NewRefreshSession(refresh, d.Clock),
		logout:            handlers.NewLogout(refresh, refreshRepo, d.Clock),
		invitationSession: handlers.NewInvitationSession(invitation, d.Clock),

		// workers
		refreshTokensCleaner: infra.NewRefreshTokensCleaner(refreshRepo, d.Log, d.Clock),
	}
}

func (m *Module) Name() string { return "auth" }

// AccessVerifier — bridge. Turns an access token into the actor it stands
// for. Used by the HTTP and websocket auth middleware.
func (m *Module) AccessVerifier() contract.Verifier[actor.Actor] { return m.accessVerifier }

// InvitationVerifier — bridge. Turns an invitation session token into the
// invitation it was opened for. Used by the invitation web pages.
func (m *Module) InvitationVerifier() contract.Verifier[contract.InvitationClaims] {
	return m.invitationVerifier
}

func (m *Module) PasswordVerifier() contract.PasswordVerifier {
	return m.passwordVerifier
}

// Register — subscribe on commands/queries/events. With error awareness
func (m *Module) Register(a *core.App) error {
	return errors.Join(
		bus.RegisterCommand(a.Commands, m.emailLogin.Handle),
		bus.RegisterCommand(a.Commands, m.refreshSession.Handle),
		bus.RegisterCommand(a.Commands, m.logout.Handle),
		bus.RegisterCommand(a.Commands, m.invitationSession.Handle),
	)
}

func (m *Module) RefreshTokensCleaner() contract.RefreshTokensCleaner { return m.refreshTokensCleaner }
