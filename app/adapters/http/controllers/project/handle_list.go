package project

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *ProjectController) List(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("workspace_id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}
	projects, err := s.projectService.List(c.Request.Context(), id)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}
	utils.Success(c, projects)
}
