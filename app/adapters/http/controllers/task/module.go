package task

import (
	"gitlab.com/shaninalex/lumna/app/services"
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = c.Provide(services.NewTaskService)
	_ = c.Provide(services.NewEntityEventService)

	_ = c.Provide(NewTaskController)
	return nil
}
