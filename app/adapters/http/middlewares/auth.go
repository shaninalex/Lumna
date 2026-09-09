package middlewares

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/pkg"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/services/auth"
)

var (
	ErrorAuthMiddlewareUnauthorized = errors.New("unauthorized")
	ErrorAuthMiddlewareInvalidID    = errors.New("invalid user id")
	ErrorIdentityNotFound           = errors.New("identity not found with given id")
)

// AuthMiddleware - dedicated type for middleware, it helps dig.Container since it based on types
type AuthMiddleware gin.HandlerFunc

func ProvideAuthMiddleware(
	identityRepository repositories.IdentityRepository,
) AuthMiddleware {
	return func(c *gin.Context) {
		accessJWTToken, err := c.Cookie("access_token")
		if err != nil {
			log.Printf("[AuthMiddleware]: %s", err)
			utils.Error(c, http.StatusUnauthorized, ErrorAuthMiddlewareUnauthorized)
			c.Abort()
			return
		}

		if accessJWTToken == "" {
			log.Printf("[AuthMiddleware]: access token \"%s\" is invalid", accessJWTToken)
			utils.Error(c, http.StatusUnauthorized, ErrorAuthMiddlewareUnauthorized)
			c.Abort()
			return
		}

		claims, err := auth.ParseAccessJWTToken(accessJWTToken)
		if err != nil {
			log.Printf("[AuthMiddleware]: unable to parse: %s", err)
			utils.Error(c, http.StatusUnauthorized, ErrorAuthMiddlewareUnauthorized)
			c.Abort()
			return
		}

		userID, err := strconv.Atoi(claims.Subject)
		if err != nil {
			utils.Error(c, http.StatusUnauthorized, ErrorAuthMiddlewareInvalidID)
			c.Abort()
			return
		}

		ctx := c.Request.Context()

		// set identity in context
		identity, err := identityRepository.GetIdentityByID(ctx, userID)
		if err != nil {
			utils.Error(c, http.StatusUnauthorized, ErrorIdentityNotFound)
			c.Abort()
			return
		}
		ctx = context.WithValue(ctx, pkg.ContextIdentity, identity)

		c.Set(pkg.ContextUserID, userID)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
