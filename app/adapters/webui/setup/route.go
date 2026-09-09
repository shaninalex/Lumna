package setup

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
	"gitlab.com/shaninalex/lumna/app/services/persistence"
	"gorm.io/gorm"
)

func RegisterSetupRoute(router *gin.Engine, conf *config.Config) {
	if !conf.Bool("serve.setup") {
		return
	}

	db := persistence.ProvideDB(conf)
	router.GET("/setup", setupIdentityMiddleware(db), handleSetup(db))
	router.POST("/setup", setupIdentityMiddleware(db), handleSetupSubmit(db))
}

func setupIdentityMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var identities []*models.Identity
		if err := db.WithContext(ctx).Find(&identities).Error; err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if len(identities) > 0 {
			ctx.Redirect(http.StatusFound, "/")
			return
		}

		ctx.Next()
	}
}
