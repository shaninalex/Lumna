package tracker

import (
	"log/slog"

	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/platform/database"
)

type Deps struct {
	DB  *database.DB
	Log *slog.Logger
}

type Module struct {
	deps Deps
}

func New(d Deps) *Module {
	return &Module{
		deps: d,
	}
}

func (m *Module) Name() string { return "tracker" }

func (m *Module) Register(a *core.App) error {
	//return errors.Join(
	//bus.RegisterCommand(a.Commands, m.command.Handle),
	//)
	return nil
}
