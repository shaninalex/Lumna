package web

import (
	"errors"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/adapters/httpx"
	"gitlab.com/shaninalex/lumna/app/adapters/web/templates"
	"gitlab.com/shaninalex/lumna/app/core"
	authContract "gitlab.com/shaninalex/lumna/app/modules/auth/contract"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
)

var ConfirmInvitationEmptyKey = errors.New("invalid invitation key")

func handleConfirmInvitation(app core.Resolve, cookieConfig httpx.CookieConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Query("key")
		if key == "" {
			c.Data(http.StatusBadRequest, "text/html; charset=utf-8", []byte(ConfirmInvitationEmptyKey.Error()))
			return
		}

		// Validation
		invitation, err := contract.AskInvitationValidate(c.Request.Context(), app(), contract.InvitationValidate{
			Hash: key,
		})
		if err != nil {
			httpx.Fail(c, errors.Join(ConfirmInvitationEmptyKey, err))
			return
		}

		// create temp guest session
		tempSession, err := authContract.ExecTempSession(c.Request.Context(), app(), authContract.TempSession{
			Email: invitation.Email,
			Role:  invitation.Role,
		})
		if err != nil {
			httpx.Fail(c, errors.Join(ConfirmInvitationEmptyKey, err))
			return
		}

		c.SetCookie(
			"invitation_access_token",
			tempSession.AccessToken,
			int(tempSession.AccessTokenDuration.Seconds()),
			"/invirarion/password", "",
			cookieConfig.Secure,
			true,
		)

		// redirect to /invitation/password
		c.Redirect(http.StatusFound, "/invitation/password")
	}
}

func confirmInvitationMiddleware(app core.Resolve) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: validate temp session
		ctx.Next()
	}
}

func handleCreatePassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		// validate temp session cookie
		// 4) html form[action=/accept-invitation] submit form to /accept-invitation
		templ.Handler(templates.ConfirmInvitationView()).ServeHTTP(c.Writer, c.Request)
	}
}

func handleAcceptInvitation() gin.HandlerFunc {
	return func(c *gin.Context) {
		// validate temp session cookie
		// request body
		// 5) Create user
		// 5.2) remove temp session cookie
		// 6) redirect to /auth/login
		c.Redirect(http.StatusFound, "/auth/login")
	}
}
