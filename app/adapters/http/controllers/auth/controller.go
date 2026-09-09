package auth

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/services/auth"
)

type AuthController struct {
	localProvider    *auth.EmailAuthProvider
	authTokenService auth.TokenService
}

func NewAuthContoller(
	authProvider *auth.EmailAuthProvider,
	authTokenService auth.TokenService,
) *AuthController {
	s := &AuthController{
		localProvider:    authProvider,
		authTokenService: authTokenService,
	}

	return s
}

func (s *AuthController) Register(router *gin.RouterGroup) {
	router.POST("/login", s.handleLogin)
	router.GET("/logout", s.handleLogout)

	// refresh requires authenticated request to get user id
	router.GET("/refresh", s.handleRefresh)

	// router.GET("/api/v1/auth/oauth/github", nil)
	// router.GET("/api/v1/auth/oauth/github/callback", nil)

	// router.GET("/api/v1/auth/oauth/google", nil)
	// router.GET("/api/v1/auth/oauth/google/callback", nil)
}
