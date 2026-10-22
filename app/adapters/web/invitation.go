package web

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/adapters/web/templates"
)

var ConfirmInvitationEmptyKey = "Confirm invitation empty key"

func handleConfirmInvitation() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Query("key")
		if key == "" {
			c.Data(http.StatusBadRequest, "text/html; charset=utf-8", []byte(ConfirmInvitationEmptyKey))
			return
		}

		// 1) Validation:
		//The invitation exists.
		//It has not expired.
		//It has not already been accepted.
		//It has not been revoked.

		// 2) create temp guest session
		// 3) redirect to /create-password

		c.Redirect(http.StatusFound, "/invirarion/password")
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
		// 6) redirect to /auth/login
		c.Redirect(http.StatusFound, "/auth/login")
	}
}
