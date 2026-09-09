package column

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *Controller) Delete(c *gin.Context) {
	columnId, err := strconv.Atoi(c.Param("column_id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	if err := s.columnService.Delete(c.Request.Context(), columnId); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	utils.Success(c, nil, "Column deleted")
}
