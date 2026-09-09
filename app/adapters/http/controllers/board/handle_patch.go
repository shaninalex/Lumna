package board

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *BoardController) Patch(c *gin.Context) {
	listId, err := strconv.Atoi(c.Param("board_id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}
	payload := struct {
		ProjectId int    `json:"project_id"`
		Title     string `json:"title"`
	}{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	list, err := s.listService.Get(c.Request.Context(), listId)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	list.Title = payload.Title

	if err := s.listService.Update(c.Request.Context(), list); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	utils.Success(c, list)
}
