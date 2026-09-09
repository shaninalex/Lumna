package project

import (
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/services"
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = c.Provide(repositories.NewGormProjectRepository)
	_ = c.Provide(services.NewProjectService)
	_ = c.Provide(services.NewListService)

	_ = c.Provide(NewProjectsController)
	return nil
}
