package column

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/services"
)

type Controller struct {
	columnService services.ColumnService
}

func NewStatusController(columnService services.ColumnService) *Controller {
	s := &Controller{
		columnService: columnService,
	}

	return s
}

func (s *Controller) Register(router *gin.RouterGroup) {
	router.GET("columns", s.List)
	router.POST("columns", s.Create)
	router.POST("columns/reorder", s.handleReorder)
	router.DELETE("columns/:column_id", s.Delete)
	router.PATCH("columns/:column_id", s.Patch)
}
