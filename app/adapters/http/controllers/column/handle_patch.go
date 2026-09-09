package column

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/dto"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *Controller) Patch(c *gin.Context) {
	columnId, err := strconv.Atoi(c.Param("column_id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	payload := dto.ColumnDTO{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	column, err := s.columnService.Get(c.Request.Context(), columnId)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	column.Title = payload.Title
	column.Position = payload.Position

	if err := s.columnService.Save(c.Request.Context(), column); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	utils.Success(c, column)
}
