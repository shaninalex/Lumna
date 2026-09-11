package api

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/adapters/api/controllers/auth"
	"gitlab.com/shaninalex/lumna/app/adapters/api/controllers/user"
	"gitlab.com/shaninalex/lumna/app/adapters/api/controllers/workspace"
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

// RegisterApiRouter mounts the API.
func RegisterApiRouter(resolve core.Resolve, verifier authc.Verifier, cfg Config, router *gin.Engine) {
	cookies := transport.CookieConfig{Secure: cfg.SecureCookies}
	router.Use(middlewares.CORSMiddleware(cfg.CORSOrigins))

	public := router.Group("/api/v1")
	auth.Register(resolve, cookies, public.Group("auth"))

	private := router.Group("/api/v1")
	private.Use(middlewares.AuthMiddleware(verifier))
	user.Register(resolve, private.Group("user"))
	workspace.Register(resolve, private.Group("workspace"))
}
