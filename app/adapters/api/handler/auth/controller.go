package auth

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/adapters/httpx"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/bus"
	"gitlab.com/shaninalex/lumna/app/modules/auth/contract"
)

func Register(resolve core.Resolve, cookies httpx.CookieConfig, router *gin.RouterGroup) {
	router.POST("/login", handleLogin(resolve, cookies))
	router.POST("/logout", handleLogout(resolve, cookies))
	router.POST("/refresh", handleRefresh(resolve, cookies))
}

func handleLogin(resolve core.Resolve, cookies httpx.CookieConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		payload := passwordLoginDTO{}
		if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
			httpx.Invalid(c, err)
			return
		}
		if err := payload.Validate(); err != nil {
			httpx.Invalid(c, err)
			return
		}

		session, err := contract.ExecEmailLogin(c.Request.Context(), resolve(), contract.EmailLogin{
			Email:    payload.Email,
			Password: bus.Secret(payload.Password),
		})
		if err != nil {
			// Status comes from the error's class: wrong credentials are 401,
			// not 400.
			httpx.Fail(c, err)
			return
		}

		writeSession(c, cookies, session)
		httpx.Success(c, nil, "Login successful")
	}
}

func handleRefresh(resolve core.Resolve, cookies httpx.CookieConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshToken, _ := c.Cookie(httpx.CookieRefreshToken)

		session, err := contract.ExecRefreshSession(c.Request.Context(), resolve(), contract.RefreshSession{
			RefreshToken: bus.Secret(refreshToken),
		})
		if err != nil {
			httpx.ClearSession(c, cookies)
			httpx.Fail(c, err)
			return
		}

		writeSession(c, cookies, session)
		httpx.Success(c, nil, "Session refreshed")
	}
}

func handleLogout(resolve core.Resolve, cookies httpx.CookieConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshToken, _ := c.Cookie(httpx.CookieRefreshToken)

		if _, err := contract.ExecLogout(c.Request.Context(), resolve(), contract.Logout{
			RefreshToken: bus.Secret(refreshToken),
		}); err != nil {
			httpx.Fail(c, err)
			return
		}

		httpx.ClearSession(c, cookies)
		httpx.Success(c, nil, "Logged out")
	}
}

func writeSession(c *gin.Context, cookies httpx.CookieConfig, s contract.SessionView) {
	httpx.SetSession(c, cookies,
		s.AccessToken, s.AccessTokenDuration,
		s.RefreshToken, s.RefreshTokenDuration,
	)
}
