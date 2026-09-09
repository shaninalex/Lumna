package invitation

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

var (
	ErrorInvitationList = errors.New("unable to get invitation list")
)

func (s *InvitationController) List(c *gin.Context) {
	invitations, err := s.invitationService.List(c.Request.Context())
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, ErrorInvitationList)
		return
	}
	utils.Success(c, invitations)
}
