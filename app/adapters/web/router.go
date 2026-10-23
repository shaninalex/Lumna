package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna"
	"gitlab.com/shaninalex/lumna/app/adapters/httpx"
	"gitlab.com/shaninalex/lumna/app/core"
	authc "gitlab.com/shaninalex/lumna/app/modules/auth/contract"
	"gitlab.com/shaninalex/lumna/app/platform/config"
)

func RegisterWebRoutes(
	cfg *config.Config,
	resolve core.Resolve,
	invitationVerifier authc.Verifier[authc.InvitationClaims],
	router *gin.Engine,
) {
	cookies := httpx.CookieConfig{Secure: cfg.SecureCookies()}

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
	router.GET("/invitation/confirm", handleConfirmInvitation(resolve, cookies))

	// Everything after /confirm runs inside the invitation session it opened.
	invitation := router.Group("/invitation", invitationSessionMiddleware(invitationVerifier))
	invitation.GET("/password", handleCreatePassword())
	invitation.POST("/accept", handleAcceptInvitation())
}
