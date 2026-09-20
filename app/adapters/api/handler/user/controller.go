package user

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/adapters/api/transport"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/actor"
	"gitlab.com/shaninalex/lumna/app/core/errs"
	"gitlab.com/shaninalex/lumna/app/modules/identity/contract"
)

func Register(resolve core.Resolve, router *gin.RouterGroup) {
	router.GET("/me", handleMe(resolve))
}

func handleMe(resolve core.Resolve) gin.HandlerFunc {
	return func(c *gin.Context) {
		// AuthMiddleware guarantees this, so a miss means the route was
		// mounted on the wrong group.
		a, ok := actor.From(c.Request.Context())
		if !ok {
			transport.Fail(c, errs.Unauthenticated("API001", "authentication required"))
			return
		}

		identity, err := contract.AskGetProfile(c.Request.Context(), resolve(), contract.GetProfile{IdentityID: a.IdentityID})
		if err != nil {
			transport.Fail(c, err)
			return
		}

		transport.Success(c, ProfileDTO{
			ID:       identity.ID,
			Email:    identity.Email,
			FullName: identity.FullName,
			Active:   identity.Active,
		})
	}
}
