package workspace

import (
	"gitlab.com/shaninalex/lumna/app/services"
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = c.Provide(services.NewWorkspaceService)

	_ = c.Provide(NewWorkspaceController)
	return nil
}
