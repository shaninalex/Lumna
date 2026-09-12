package auth

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/adapters/api/transport"
	"gitlab.com/shaninalex/lumna/app/adapters/api/validators"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/bus"
	"gitlab.com/shaninalex/lumna/app/modules/auth/contract"
)

func Register(resolve core.Resolve, cookies transport.CookieConfig, router *gin.RouterGroup) {
	router.POST("/login", handleLogin(resolve, cookies))
	router.POST("/logout", handleLogout(resolve, cookies))
	router.POST("/refresh", handleRefresh(resolve, cookies))
}

func handleLogin(resolve core.Resolve, cookies transport.CookieConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		payload := PasswordCredentials{}
		if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
			transport.Invalid(c, err)
			return
		}
		if err := validators.Validate(payload); err != nil {
			transport.Invalid(c, err)
			return
		}

		session, err := contract.ExecEmailLogin(c.Request.Context(), resolve(), contract.EmailLogin{
			Email:    payload.Email,
			Password: bus.Secret(payload.Password),
		})
		if err != nil {
			// Status comes from the error's class: wrong credentials are 401,
			// not 400.
			transport.Fail(c, err)
			return
		}

		writeSession(c, cookies, session)
		transport.Success(c, nil, "Login successful")
	}
}

func handleRefresh(resolve core.Resolve, cookies transport.CookieConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshToken, _ := c.Cookie(transport.CookieRefreshToken)

		session, err := contract.ExecRefreshSession(c.Request.Context(), resolve(), contract.RefreshSession{
			RefreshToken: bus.Secret(refreshToken),
		})
		if err != nil {
			transport.ClearSession(c, cookies)
			transport.Fail(c, err)
			return
		}

		writeSession(c, cookies, session)
		transport.Success(c, nil, "Session refreshed")
	}
}

func handleLogout(resolve core.Resolve, cookies transport.CookieConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshToken, _ := c.Cookie(transport.CookieRefreshToken)

		if _, err := contract.ExecLogout(c.Request.Context(), resolve(), contract.Logout{
			RefreshToken: bus.Secret(refreshToken),
		}); err != nil {
			transport.Fail(c, err)
			return
		}

		transport.ClearSession(c, cookies)
		transport.Success(c, nil, "Logged out")
	}
}

func writeSession(c *gin.Context, cookies transport.CookieConfig, s contract.SessionView) {
	transport.SetSession(c, cookies,
		s.AccessToken, s.AccessTokenDuration,
		s.RefreshToken, s.RefreshTokenDuration,
	)
}
