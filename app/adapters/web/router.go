package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/platform/config"
)

func RegisterWebRoutes(cfg *config.Config, resolve core.Resolve, router *gin.Engine) {
	if cfg.SetupEnabled() {
		router.GET("/setup", setupIdentityMiddleware(resolve), handleSetup())
		router.POST("/setup", setupIdentityMiddleware(resolve), handleSetupSubmit(resolve))
	}

	if cfg.EmbedSPA() {
		RegisterSPA(router)
	}

	docsFS := lumna.StaticFS("resources/openapi")
	if docsFS != nil {
		router.StaticFS("/api-docs", http.FS(docsFS))
	}
	router.GET("/invitation/confirm", handleConfirmInvitation())
	router.GET("/invitation/password", handleCreatePassword())
	router.POST("/invitation/accept", handleAcceptInvitation())

}
