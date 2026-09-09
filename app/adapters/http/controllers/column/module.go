package column

import (
	"gitlab.com/shaninalex/lumna/app/services"
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = c.Provide(services.NewStatusService)

	_ = c.Provide(NewStatusController)
	return nil
}
