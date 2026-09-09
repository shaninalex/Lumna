package workspace

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

type routeWorkspaceListParams struct {
	Active *bool `form:"active,omitempty"`
}

func (s *WorkspaceController) handlerList(c *gin.Context) {
	var query routeWorkspaceListParams
	err := c.ShouldBindQuery(&query)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err)
		return
	}

	params := make(map[string]any)
	if query.Active != nil {
		params["active"] = *query.Active
	}

	workspaces, err := s.workspaceManager.List(c.Request.Context(), params)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err)
		return
	}

	utils.Success(c, workspaces)
}
