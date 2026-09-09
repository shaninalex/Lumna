package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/models"
	pkgUtils "gitlab.com/shaninalex/lumna/app/pkg/utils"
)

func (s *ProjectController) Create(c *gin.Context) {
	var payload models.ProjectCreateModel
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	userID, err := utils.GetUserID(c)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	payload.OwnerID = pkgUtils.Pointer(userID)
	project, err := s.projectService.Create(c.Request.Context(), payload)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	utils.Success(c, project)
}
