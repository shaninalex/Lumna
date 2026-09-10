package bootstrap

import (
	"context"
	"fmt"
	"io/fs"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/bus"
	"gitlab.com/shaninalex/lumna/app/platform/config"
	"gitlab.com/shaninalex/lumna/app/platform/database"
)

type App struct {
	Core    *core.App
	HTTP    *gin.Engine
	Log     *slog.Logger
	modules []core.Module
}

type Assets struct {
	Migrations fs.FS
	OpenAPI    fs.FS
	SPA        fs.FS
	Version    string
}

func New(ctx context.Context, cfg *config.Config, assets Assets) (*App, error) {
	// ======= Platform (logger, db...) =======
	log := slog.New(slog.NewTextHandler(os.Stderr, nil))
	db := database.New(cfg)

	// ======= Core =======
	write := []bus.Middleware{
		// pmw.Recover(log),
		// pmw.RequestLog(log),
		// pmw.Authorize(guard), // permisions BEFORE open transactions
		// pmw.Transaction(db),  // tx in ctx; commit/rollback
		// pmw.Metrics(),
	}

	read := []bus.Middleware{
		// pmw.Recover(log),
		// pmw.RequestLog(log),
		// pmw.Authorize(guard),
		// pmw.ReadOnly(db), // Connection for read only
		// pmw.Metrics(),
	}

	c := &core.App{
		Commands: bus.NewCommandBus(write...),
		Queries:  bus.NewQueryBus(read...),
		Events:   bus.NewEventBus(),
	}

	// ======= Modules =======
	mods, err := buildModules(cfg, db, log)
	if err != nil {
		return nil, err
	}
	for _, m := range mods {
		if err := m.Register(c); err != nil {
			return nil, fmt.Errorf("register module %s: %w", m.Name(), err)
		}
	}

	// ======= Stop registering commands and queries =======
	c.Commands.Seal()
	c.Queries.Seal()

	// ======= Adapters =======
	router := gin.Default()

	return &App{
		Core:    c,
		HTTP:    router,
		Log:     log,
		modules: mods,
	}, nil
}
