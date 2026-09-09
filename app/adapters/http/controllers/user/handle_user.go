package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *UserController) Me(c *gin.Context) {
	userId, err := utils.GetUserID(c)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, fmt.Errorf("user not found in context"))
		return
	}
	identity, err := s.userService.Identity(c.Request.Context(), userId)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, fmt.Errorf("user not found"))
		return
	}

	utils.Success(c, identity)
}
