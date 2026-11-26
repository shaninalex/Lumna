package httpx

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const codeCrossOrigin = "CROSS_ORIGIN_REQUEST"

// CrossOriginMiddleware rejects unsafe requests (POST, PUT, PATCH, DELETE)
// that a browser sent on behalf of another site. It is the token-less CSRF
// defence from net/http: Sec-Fetch-Site must be same-origin (or none, for
// requests the user typed), falling back to comparing Origin with Host in
// browsers that don't send Sec-Fetch-Site. Requests with neither header are
// not from a browser and pass. No other origin is trusted: the SPA is always
// served from this one.
//
// Safe methods always pass, so it can sit in front of every route, including
// the server-rendered forms.
func CrossOriginMiddleware() gin.HandlerFunc {
	protection := http.NewCrossOriginProtection()

	return func(c *gin.Context) {
		if err := protection.Check(c.Request); err != nil {
			ReturnJSON(c, http.StatusForbidden, nil, NewApiError("cross-origin request rejected", codeCrossOrigin, nil))
			c.Abort()
			return
		}
		c.Next()
	}
}
