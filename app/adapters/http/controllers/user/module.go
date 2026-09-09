package user

import (
	"gitlab.com/shaninalex/lumna/app/services/user"
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = c.Provide(user.NewUserService)

	_ = c.Provide(NewUserController)
	return nil
}
