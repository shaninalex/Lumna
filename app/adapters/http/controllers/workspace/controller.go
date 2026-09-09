package workspace

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/services"
)

type WorkspaceController struct {
	workspaceManager services.WorkspaceManager
}

func NewWorkspaceController(workspaceManager services.WorkspaceManager) *WorkspaceController {
	s := &WorkspaceController{
		workspaceManager: workspaceManager,
	}

	return s
}

func (s *WorkspaceController) Register(router *gin.RouterGroup) {
	router.GET("workspaces", s.handlerList)
	router.POST("workspaces", s.handlerCreate)
}
