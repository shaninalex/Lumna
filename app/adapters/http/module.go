package api

import (
	"gitlab.com/shaninalex/lumna/app/adapters/http/controllers/auth"
	"gitlab.com/shaninalex/lumna/app/adapters/http/controllers/board"
	"gitlab.com/shaninalex/lumna/app/adapters/http/controllers/column"
	"gitlab.com/shaninalex/lumna/app/adapters/http/controllers/invitation"
	"gitlab.com/shaninalex/lumna/app/adapters/http/controllers/project"
	"gitlab.com/shaninalex/lumna/app/adapters/http/controllers/task"
	"gitlab.com/shaninalex/lumna/app/adapters/http/controllers/user"
	"gitlab.com/shaninalex/lumna/app/adapters/http/controllers/workspace"
	"gitlab.com/shaninalex/lumna/app/adapters/http/middlewares"
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = auth.Module(c)
	_ = board.Module(c)
	_ = column.Module(c)
	_ = project.Module(c)
	_ = task.Module(c)
	_ = user.Module(c)
	_ = invitation.Module(c)
	_ = workspace.Module(c)
	_ = middlewares.Module(c)

	_ = c.Provide(NewApi)

	return nil
}
