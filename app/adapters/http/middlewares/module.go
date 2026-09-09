package middlewares

import (
	"gitlab.com/shaninalex/lumna/app/repositories"
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = c.Provide(repositories.NewGormIdentityRepository)

	_ = c.Provide(ProvideAuthMiddleware)

	return nil
}
