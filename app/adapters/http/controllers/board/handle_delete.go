package board

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *BoardController) Delete(c *gin.Context) {
	listId, err := strconv.Atoi(c.Param("board_id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	if err := s.listService.Delete(c.Request.Context(), listId); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	utils.Success(c, nil, "Board deleted")
}
