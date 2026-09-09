package workspace

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/models"
)

func (s *WorkspaceController) handlerCreate(c *gin.Context) {
	var data models.WorkspaceCreateModel
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	if err := data.Validate(); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	identity, err := utils.GetIdentity(c.Request.Context())
	if err != nil {
		utils.Error(c, http.StatusUnauthorized, err)
		return
	}

	workspace, err := s.workspaceManager.CreateWithOwner(c.Request.Context(), data.Title, identity)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err)
		return
	}

	utils.Success(c, workspace)
}
