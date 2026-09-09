package column

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/services"
)

func (s *Controller) Create(c *gin.Context) {
	payload := services.BoardCreatePayload{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	column := &models.Column{
		Title:    payload.Title,
		BoardId:  payload.BoardId,
		Position: payload.Order,
	}

	err := s.columnService.Save(c.Request.Context(), column)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	utils.Success(c, column)
}
