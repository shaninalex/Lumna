package invitation

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/services/invitation"
)

type InvitationController struct {
	invitationService invitation.Service
}

func NewInvitationController(
	invitationService invitation.Service,
) *InvitationController {
	s := &InvitationController{
		invitationService: invitationService,
	}

	return s
}

func (s *InvitationController) Register(router *gin.RouterGroup) {
	router.GET("invitations", s.List)
	router.POST("invitations", s.Create)
}
