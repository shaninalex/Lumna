package column

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/dto"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *Controller) handleReorder(c *gin.Context) {
	var data dto.ColumnReorderingDto
	if err := c.BindJSON(&data); err != nil {
		utils.Error(c, 400, err)
		return
	}

	if err := s.columnService.Reorder(c.Request.Context(), data.ColumnsOrder); err != nil {
		utils.Error(c, 400, err)
		return
	}

	columns := s.columnService.Filter(c.Request.Context(), data.BoardId)
	utils.Success(c, columns)
}
