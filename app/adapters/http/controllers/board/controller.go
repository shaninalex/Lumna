package board

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/services"
)

type BoardController struct {
	listService   services.BoardService
	columnService services.ColumnService
	taskService   services.TaskService
}

func NewListController(
	listService services.BoardService,
	columnService services.ColumnService,
	taskService services.TaskService,
) *BoardController {
	s := &BoardController{
		listService:   listService,
		columnService: columnService,
		taskService:   taskService,
	}
	return s
}

func (s *BoardController) Register(router *gin.RouterGroup) {
	router.POST("boards", s.Create)
	router.GET("boards", s.List)
	router.GET("boards/:board_id", s.Get)
	router.PATCH("boards/:board_id", s.Patch)
	router.DELETE("boards/:board_id", s.Delete)
	router.PATCH("boards/:board_id/order", s.ChangeOrder)
}
