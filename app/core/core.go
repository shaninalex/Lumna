package core

import "gitlab.com/shaninalex/lumna/app/core/bus"

type App struct {
	Commands *bus.CommandBus
	Queries  *bus.QueryBus
	Events   *bus.EventBus
}

type Module interface {
	Name() string
	Register(a *App) error
}
