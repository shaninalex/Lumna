package user

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/services/invitation"
	"gitlab.com/shaninalex/lumna/app/services/user"
)

type UserController struct {
	userService       user.Service
	invitationManager invitation.Service
}

func NewUserController(
	userService user.Service,
	invitationManager invitation.Service,
) *UserController {
	s := &UserController{
		userService:       userService,
		invitationManager: invitationManager,
	}

	return s
}

func (s *UserController) Register(router *gin.RouterGroup) {
	router.GET("/user/me", s.Me)
}
