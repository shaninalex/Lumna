package auth

import (
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/services/auth"
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	if err := c.Provide(repositories.NewGormCredentialRepository); err != nil {
		panic(err)
	}
	if err := c.Provide(repositories.NewGormIdentityRepository); err != nil {
		panic(err)
	}
	if err := c.Provide(auth.NewEmailAuthProvider); err != nil {
		panic(err)
	}
	if err := c.Provide(repositories.NewGormRefreshTokenRepository); err != nil {
		panic(err)
	}
	if err := c.Provide(auth.NewAuthTokenService); err != nil {
		panic(err)
	}

	if err := c.Provide(NewAuthContoller); err != nil {
		panic(err)
	}
	return nil
}
