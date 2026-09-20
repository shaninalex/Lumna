package bootstrap

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"

	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/bus"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
	"gitlab.com/shaninalex/lumna/app/platform/config"
	"gitlab.com/shaninalex/lumna/app/platform/database"
	pmw "gitlab.com/shaninalex/lumna/app/platform/middleware"
)

type Instance struct {
	Core    *core.App
	Log     *slog.Logger
	Clock   clock.Clock
	Bridges Bridges
	modules []core.Module

	closers []func(context.Context) error
}

// Close clear resources. Required for CLI
func (a *Instance) Close(ctx context.Context) error {
	var errs []error
	for i := len(a.closers) - 1; i >= 0; i-- { // backwards
		if err := a.closers[i](ctx); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

type Assets struct {
	Migrations fs.FS
	OpenAPI    fs.FS
	SPA        fs.FS
	Version    string
}

func New(ctx context.Context, cfg *config.Config, assets Assets) (*Instance, error) {
	// ======= Platform (logger, db, clock...) =======
	log := slog.New(slog.NewTextHandler(os.Stderr, nil))
	db := database.New(cfg)
	clk := clock.System()

	// ======= Core =======
	write := []bus.Middleware{
		// pmw.Recover(log),
		// pmw.RequestLog(log),
		// pmw.Authorize(guard), // permissions BEFORE open transactions
		pmw.Transaction(db), // tx in ctx; commit/rollback
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
	mods, bridges, err := buildModules(cfg, db, log, clk, c.Events)
	if err != nil {
		return nil, err
	}
	for _, m := range mods {
		if err := m.Register(c); err != nil {
			return nil, fmt.Errorf("register module %s: %w", m.Name(), err)
		}
	}

	// ======= Stop registering commands, queries and events =======
	c.Commands.Seal()
	c.Queries.Seal()
	c.Events.Seal()

	return &Instance{
		Core:    c,
		Log:     log,
		Clock:   clk,
		Bridges: bridges,
		modules: mods,
	}, nil
}
