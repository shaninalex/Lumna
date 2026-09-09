package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *AuthController) handleLogout(c *gin.Context) {
	userID, _ := utils.GetUserID(c)

	if err := s.authTokenService.DeleteRefreshToken(c.Request.Context(), userID); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	c.SetCookie("access_token", "", -1, "/", "", true, true)
	c.SetCookie("refresh_token", "", -1, "/", "", true, true)

	utils.Success(c, nil, "Logout successfull")
}
