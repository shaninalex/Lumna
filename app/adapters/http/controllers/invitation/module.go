package invitation

import (
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/services/invitation"
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = c.Provide(repositories.NewGormInvitationRepository)
	_ = c.Provide(invitation.ProvideInvitationService)

	if err := c.Provide(NewInvitationController); err != nil {
		return err
	}

	return nil
}
