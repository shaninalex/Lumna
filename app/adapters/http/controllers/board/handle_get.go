package board

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *BoardController) Get(c *gin.Context) {
	listId, err := strconv.Atoi(c.Param("board_id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	list, err := s.listService.Get(c.Request.Context(), listId)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	utils.Success(c, list)
}
