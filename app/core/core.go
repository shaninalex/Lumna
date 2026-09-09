package core

type App struct {
	// Commands *bus.CommandBus
	// Queries  *bus.QueryBus
	// Events   *bus.EventBus
}

type Module interface {
	Name() string
	Register(a *App) error
}
