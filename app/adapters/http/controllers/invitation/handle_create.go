package invitation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

type InvitationCreatePayload struct {
	Email string `json:"email" binding:"required"`
	Role  string `json:"role" binding:"required"`
}

func (s *InvitationController) Create(c *gin.Context) {
	payload := InvitationCreatePayload{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	invitation, _, err := s.invitationService.Create(c.Request.Context(), payload.Email, payload.Role, nil)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
	}

	utils.Success(c, invitation)
}
