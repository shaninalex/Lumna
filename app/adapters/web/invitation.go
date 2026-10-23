package web

import (
	"errors"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/adapters/httpx"
	"gitlab.com/shaninalex/lumna/app/adapters/web/templates"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/actor"
	authContract "gitlab.com/shaninalex/lumna/app/modules/auth/contract"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
)

var ConfirmInvitationEmptyKey = errors.New("invalid invitation key")
var TempSessionCookieAccessToken = "invitation_access_token"

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
			TempSessionCookieAccessToken,
			tempSession.AccessToken,
			int(tempSession.AccessTokenDuration.Seconds()),
			"/invitation", "",
			cookieConfig.Secure,
			true,
		)

		// redirect to /invitation/password
		c.Redirect(http.StatusFound, "/invitation/password")
	}
}

func confirmInvitationMiddleware(verifier authContract.Verifier) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie(TempSessionCookieAccessToken)
		if err != nil {
			log.Println(err)
			c.Data(http.StatusForbidden, "text/html; charset=utf-8", []byte(ConfirmInvitationEmptyKey.Error()))
			return
		}

		a, err := verifier.Verify(c.Request.Context(), accessToken)
		if err != nil {
			log.Println(err)
			c.Data(http.StatusForbidden, "text/html; charset=utf-8", []byte(ConfirmInvitationEmptyKey.Error()))
			return
		}

		c.Request = c.Request.WithContext(actor.With(c.Request.Context(), a))
		c.Next()
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
