package setup

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/web/setup/templates"
	"gorm.io/gorm"
)

func handleSetup(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		templ.Handler(templates.SetupView(templates.SetupViewData{})).ServeHTTP(c.Writer, c.Request)
	}
}
