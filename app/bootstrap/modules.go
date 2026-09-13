package bootstrap

import (
	"errors"
	"log/slog"
	"time"

	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/modules/auth"
	authc "gitlab.com/shaninalex/lumna/app/modules/auth/contract"
	"gitlab.com/shaninalex/lumna/app/modules/identity"
	"gitlab.com/shaninalex/lumna/app/modules/tracker"
	"gitlab.com/shaninalex/lumna/app/modules/workspace"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
	"gitlab.com/shaninalex/lumna/app/platform/config"
	"gitlab.com/shaninalex/lumna/app/platform/database"
)

const (
	defaultAccessTTL  = 15 * time.Minute
	defaultRefreshTTL = 30 * 24 * time.Hour
	tokenIssuer       = "lumna"
)

// Bridges are module-provided ports that an adapter needs directly, without
// going through the bus. It's basically a "library call" without side effects
type Bridges struct {
	AuthVerifier authc.Verifier
}

func buildModules(cfg *config.Config, db *database.DB, log *slog.Logger, clk clock.Clock) ([]core.Module, Bridges, error) {
	secret := []byte(cfg.AuthSecret())
	if len(secret) == 0 {
		return nil, Bridges{}, errors.New("bootstrap: secret_key is empty, cannot sign tokens")
	}

	identityModule := identity.New(identity.Deps{
		DB:     db,
		Log:    log,
		Clock:  clk,
		Secret: secret,
	})

	authModule := auth.New(auth.Deps{
		DB:         db,
		Log:        log,
		Clock:      clk,
		Secret:     secret,
		AccessTTL:  minutesOr(cfg.Int("auth.access_ttl"), defaultAccessTTL),
		RefreshTTL: minutesOr(cfg.Int("auth.refresh_ttl"), defaultRefreshTTL),
		Issuer:     tokenIssuer,

		Identities:  identityModule.Authenticator(),
		Provisioner: identityModule.Provisioner(),
	})

	trackerModule := tracker.New(tracker.Deps{DB: db, Log: log, Clock: clk})
	workspaceModule := workspace.New(workspace.Deps{DB: db, Log: log, Clock: clk})

	modules := []core.Module{
		identityModule,
		authModule,
		trackerModule,
		workspaceModule,
	}

	bridges := Bridges{
		AuthVerifier: authModule.Verifier(),
	}

	return modules, bridges, nil
}

func minutesOr(minutes int, fallback time.Duration) time.Duration {
	if minutes <= 0 {
		return fallback
	}
	return time.Duration(minutes) * time.Minute
}
