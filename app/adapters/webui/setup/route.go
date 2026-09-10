package setup

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/platform/config"
)

func RegisterSetupRoute(router *gin.Engine, conf *config.Config) {
	if !conf.Bool("serve.setup") {
		return
	}

	// db := persistence.ProvideDB(conf)
	// router.GET("/setup", setupIdentityMiddleware(db), handleSetup(db))
	// router.POST("/setup", setupIdentityMiddleware(db), handleSetupSubmit(db))
	router.GET("/setup", nil)
	router.POST("/setup", nil)
}

// func setupIdentityMiddleware(db *gorm.DB) gin.HandlerFunc {
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
