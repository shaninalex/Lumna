package project

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *ProjectController) Patch(c *gin.Context) {
	payload := struct {
		Title string `json:"title"`
	}{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	project, err := s.projectService.Get(c.Request.Context(), id)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}
	project.Title = payload.Title
	if err = s.projectService.Update(c.Request.Context(), project); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	utils.Success(c, project)
}
