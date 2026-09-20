package middlewares

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/adapters/api/transport"
	"gitlab.com/shaninalex/lumna/app/core/actor"
	"gitlab.com/shaninalex/lumna/app/modules/auth/contract"
)

// AuthMiddleware turns the access-token cookie into an actor in the request
// context. It calls the auth module's Verifier directly — a signature check
// needs no bus, no transaction and no database.
//
// Everything downstream reads the actor from ctx, never a user object: that
// is what keeps handlers from depending on the identity module.
func AuthMiddleware(verifier contract.Verifier) gin.HandlerFunc {
	return func(c *gin.Context) {
		// A missing cookie is an empty token, and the verifier already
		// rejects that — one source of truth for "not authenticated".
		accessToken, _ := c.Cookie(transport.CookieAccessToken)

		a, err := verifier.Verify(c.Request.Context(), accessToken)
		if err != nil {
			transport.Fail(c, err)
			c.Abort()
			return
		}

		c.Request = c.Request.WithContext(actor.With(c.Request.Context(), a))
		c.Next()
	}
}
