package board

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *BoardController) Create(c *gin.Context) {
	payload := struct {
		ProjectId int    `json:"project_id"`
		Title     string `json:"title"`
	}{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	list, err := s.listService.Create(c.Request.Context(), payload.ProjectId, payload.Title)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
	}

	utils.Success(c, list)
}
