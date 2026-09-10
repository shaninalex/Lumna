package setup

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/adapters/webui/setup/templates"
	"gitlab.com/shaninalex/lumna/app/core"
)

func RegisterSetupRoute(app core.Resolve, router *gin.Engine) {
	router.GET("/setup", handleSetup())
	router.POST("/setup", handleSetupSubmit(app))
}

func handleSetup() gin.HandlerFunc {
	return func(c *gin.Context) {
		templ.Handler(templates.SetupView(templates.SetupViewData{})).ServeHTTP(c.Writer, c.Request)
	}
}

// Note: Legacy...
// func setupIdentityMiddleware(app core.Resolve) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var identities []*models.Identity
// 		if err := db.WithContext(ctx).Find(&identities).Error; err != nil {
// 			ctx.AbortWithStatus(http.StatusInternalServerError)
// 			return
// 		}
// 		if len(identities) > 0 {
// 			ctx.Redirect(http.StatusFound, "/")
// 			return
// 		}
// 		ctx.Next()
// 	}
// }
