package auth

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/services/auth"
)

var (
	ErrorAuthRefreshToken = errors.New("invalid refresh token")
)

func (s *AuthController) handleRefresh(c *gin.Context) {
	refreshCookie, err := c.Cookie("refresh_token")
	if err != nil {
		log.Println("[handleRefresh]: get refresh_token cookie error - ", err)
		utils.Error(c, http.StatusUnauthorized, ErrorAuthRefreshToken)
		return
	}

	if refreshCookie == "" {
		utils.Error(c, http.StatusUnauthorized, ErrorAuthRefreshToken)
		return
	}

	dbRefreshToken, err := s.authTokenService.GetByHash(c.Request.Context(), auth.ToHashToken(refreshCookie))
	if dbRefreshToken == nil {
		utils.Error(c, http.StatusBadRequest, ErrorAuthRefreshToken)
		return
	}

	// Create new set of access/refresh tokens
	accessTtl := time.Minute * 15
	token, err := auth.GenerateAccessJWTToken(dbRefreshToken.IdentityID, "all", accessTtl)
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
		IdentityID: dbRefreshToken.IdentityID,
		Hash:       toDB,
		ClientID:   "angular-web-app",
		Scopes:     "all",
		ExpiresAt:  refreshExp,
	}

	if err := s.authTokenService.CreateRefreshToken(c.Request.Context(), rt.IdentityID, &rt); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	// assign new cookies
	c.SetCookie("access_token", token, int(accessTtl.Seconds()), "/", "", true, true)
	c.SetCookie("refresh_token", toClient, int(refreshTtl.Seconds()), "/", "", true, true)

	utils.Success(c, nil)
}
