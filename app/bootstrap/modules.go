package bootstrap

import (
	"errors"
	"log/slog"
	"time"

	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/modules/auth"
	"gitlab.com/shaninalex/lumna/app/modules/identity"
	"gitlab.com/shaninalex/lumna/app/platform/config"
	"gitlab.com/shaninalex/lumna/app/platform/database"
)

const (
	defaultAccessTTL  = 15 * time.Minute
	defaultRefreshTTL = 30 * 24 * time.Hour
	tokenIssuer       = "lumna"
)

func buildModules(cfg *config.Config, db *database.DB, log *slog.Logger) ([]core.Module, error) {
	secret := []byte(cfg.AuthSecret())
	if len(secret) == 0 {
		return nil, errors.New("bootstrap: secret_key is empty, cannot sign tokens")
	}

	identityModule := identity.New(identity.Deps{
		DB:     db,
		Log:    log,
		Secret: secret,
	})

	authModule := auth.New(auth.Deps{
		DB:         db,
		Log:        log,
		Secret:     secret,
		AccessTTL:  minutesOr(cfg.Int("auth.access_ttl"), defaultAccessTTL),
		RefreshTTL: minutesOr(cfg.Int("auth.refresh_ttl"), defaultRefreshTTL),
		Issuer:     tokenIssuer,

		Identities:  identityModule.Authenticator(),
		Provisioner: identityModule.Provisioner(),
	})

	return []core.Module{
		identityModule,
		authModule,
	}, nil
}

func minutesOr(minutes int, fallback time.Duration) time.Duration {
	if minutes <= 0 {
		return fallback
	}
	return time.Duration(minutes) * time.Minute
}
