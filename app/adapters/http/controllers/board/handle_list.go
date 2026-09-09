package board

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *BoardController) List(c *gin.Context) {
	projectId, err := strconv.Atoi(c.Query("project_id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	list, err := s.listService.List(c.Request.Context(), projectId)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	utils.Success(c, list)
}
