package bootstrap

import (
	"log/slog"

	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/modules/identity"
	"gitlab.com/shaninalex/lumna/app/platform/config"
	"gitlab.com/shaninalex/lumna/app/platform/database"
)

func buildModules(cfg *config.Config, db *database.DB, log *slog.Logger) ([]core.Module, error) {
	identityModule := identity.New(identity.Deps{
		DB:     db,
		Log:    log,
		Secret: []byte(cfg.AuthSecret()),
		// Mailer: notify.Mailer(), // ← брідж
	})

	return []core.Module{
		identityModule,
	}, nil
}
