package board

import (
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/services"
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = c.Provide(repositories.NewGormBoardRepository)
	_ = c.Provide(repositories.NewGormTaskRepository)
	_ = c.Provide(repositories.NewGormColumnRepository)
	_ = c.Provide(services.NewListService)
	_ = c.Provide(services.NewStatusService)
	_ = c.Provide(services.NewTaskService)

	_ = c.Provide(NewListController)
	return nil
}
