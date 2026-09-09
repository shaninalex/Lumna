package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/services/auth"
)

func (s *AuthController) handleLogin(c *gin.Context) {
	payload := auth.PasswordCredentials{}
	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	if err := payload.Validate(); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	identity, err := s.localProvider.Authenticate(c.Request.Context(), &payload)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	accessTtl := time.Minute * 15
	token, err := auth.GenerateAccessJWTToken(identity.ID, "all", accessTtl)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	toClient, toDB, err := auth.GenerateRefreshToken()
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	refreshTtl := time.Hour * 24 * 30 // 30 days
	refreshExp := time.Now().Add(refreshTtl)
	rt := models.RefreshToken{
		IdentityID: identity.ID,
		Hash:       toDB,
		ClientID:   "angular-web-app",
		Scopes:     "all",
		ExpiresAt:  refreshExp,
	}

	if err := s.authTokenService.CreateRefreshToken(c.Request.Context(), identity.ID, &rt); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	c.SetCookie("access_token", token, int(accessTtl.Seconds()), "/", "", true, true)
	c.SetCookie("refresh_token", toClient, int(refreshTtl.Seconds()), "/", "", true, true)

	utils.Success(c, identity, "Login successfull")
}
