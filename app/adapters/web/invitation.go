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
	authContract "gitlab.com/shaninalex/lumna/app/modules/auth/contract"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
)

var ConfirmInvitationEmptyKey = errors.New("invalid invitation key")

// InvitationCookieAccessToken holds the invitation session token. It is
// scoped to /invitation, so it never reaches the API or the SPA.
var InvitationCookieAccessToken = "invitation_access_token"

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

		// open an invitation session: the token carries the invitation
		// validated above, so later steps don't need the key again
		session, err := authContract.ExecInvitationSession(c.Request.Context(), app(), authContract.InvitationSession{
			Email: invitation.Email,
			Role:  invitation.Role,
		})
		if err != nil {
			httpx.Fail(c, errors.Join(ConfirmInvitationEmptyKey, err))
			return
		}

		c.SetCookie(
			InvitationCookieAccessToken,
			session.AccessToken,
			int(session.AccessTokenDuration.Seconds()),
			"/invitation", "",
			cookieConfig.Secure,
			true,
		)

		// redirect to /invitation/password
		c.Redirect(http.StatusFound, "/invitation/password")
	}
}

// invitationClaimsKey is the gin context key the verified invitation is
// stored under by invitationSessionMiddleware.
const invitationClaimsKey = "invitation_claims"

// invitationSessionMiddleware lets a request through only with a valid
// invitation session cookie, and stores the invitation it carries for the
// handlers (see invitationFrom). A regular access token is not accepted here.
func invitationSessionMiddleware(verifier authContract.Verifier[authContract.InvitationClaims]) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie(InvitationCookieAccessToken)
		if err != nil {
			log.Println(err)
			c.Data(http.StatusForbidden, "text/html; charset=utf-8", []byte(ConfirmInvitationEmptyKey.Error()))
			c.Abort()
			return
		}

		claims, err := verifier.Verify(c.Request.Context(), accessToken)
		if err != nil {
			log.Println(err)
			c.Data(http.StatusForbidden, "text/html; charset=utf-8", []byte(ConfirmInvitationEmptyKey.Error()))
			c.Abort()
			return
		}

		c.Set(invitationClaimsKey, claims)
		c.Next()
	}
}

// invitationFrom returns the invitation of the current session. It is only
// meaningful behind invitationSessionMiddleware.
func invitationFrom(c *gin.Context) (authContract.InvitationClaims, bool) {
	v, ok := c.Get(invitationClaimsKey)
	if !ok {
		return authContract.InvitationClaims{}, false
	}
	claims, ok := v.(authContract.InvitationClaims)
	return claims, ok
}

func handleCreatePassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 4) html form[action=/accept-invitation] submit form to /accept-invitation
		templ.Handler(templates.ConfirmInvitationView()).ServeHTTP(c.Writer, c.Request)
	}
}

func handleAcceptInvitation() gin.HandlerFunc {
	return func(c *gin.Context) {
		// invitation, _ := invitationFrom(c) — email and role to create the user with
		// request body
		// 5) Create user (workspace validates the role)
		// 5.2) remove temp session cookie
		// 6) redirect to /auth/login
		c.Redirect(http.StatusFound, "/auth/login")
	}
}
