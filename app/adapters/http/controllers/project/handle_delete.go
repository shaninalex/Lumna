package project

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *ProjectController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}
	if err := s.projectService.Delete(c.Request.Context(), id); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}
	utils.Success(c, nil)
}
