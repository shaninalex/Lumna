package http

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/adapters/api/handler/auth"
	"gitlab.com/shaninalex/lumna/app/adapters/api/handler/board"
	"gitlab.com/shaninalex/lumna/app/adapters/api/handler/column"
	"gitlab.com/shaninalex/lumna/app/adapters/api/handler/project"
	"gitlab.com/shaninalex/lumna/app/adapters/api/handler/task"
	"gitlab.com/shaninalex/lumna/app/adapters/api/handler/user"
	"gitlab.com/shaninalex/lumna/app/adapters/api/handler/workspace"
	"gitlab.com/shaninalex/lumna/app/adapters/api/middlewares"
	"gitlab.com/shaninalex/lumna/app/adapters/api/transport"
	"gitlab.com/shaninalex/lumna/app/core"
	authc "gitlab.com/shaninalex/lumna/app/modules/auth/contract"
)

// Config is the deployment-dependent part of the HTTP API, resolved in
// bootstrap so no controller reaches for the config itself.
type Config struct {
	CORSOrigins   []string
	SecureCookies bool
}

// RegisterApiRoutes mounts the API.
func RegisterApiRoutes(resolve core.Resolve, verifier authc.Verifier, cfg Config, router *gin.Engine) {
	cookies := transport.CookieConfig{Secure: cfg.SecureCookies}
	router.Use(middlewares.CORSMiddleware(cfg.CORSOrigins))

	public := router.Group("/api/v1")
	auth.Register(resolve, cookies, public.Group("auth"))

	private := router.Group("/api/v1")
	private.Use(middlewares.AuthMiddleware(verifier))
	user.Register(resolve, private.Group("user"))
	workspace.Register(resolve, private.Group("workspaces"))
	project.Register(resolve, private.Group("projects"))
	board.Register(resolve, private.Group("boards"))
	column.Register(resolve, private.Group("columns"))
	task.Register(resolve, private.Group("tasks"))
}
