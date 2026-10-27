package session

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/adapters/httpx"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/actor"
	"gitlab.com/shaninalex/lumna/app/core/errs"
	"gitlab.com/shaninalex/lumna/app/modules/auth/contract"
)

func Register(resolve core.Resolve, router *gin.RouterGroup) {
	router.GET("refresh-tokens", handleRefreshTokenList(resolve))
	router.DELETE("refresh-tokens/:id", handleRefreshTokenDelete(resolve))
	router.PATCH("refresh-tokens/:id/revoke", handleRefreshTokenRevoke(resolve))
}

var (
	ErrUnauthenticatedRequest = errs.Unauthenticated("SES01", "unauthenticated request")
)

func handleRefreshTokenList(resolve core.Resolve) gin.HandlerFunc {
	return func(c *gin.Context) {
		a, ok := actor.From(c.Request.Context())
		if !ok || a.IdentityID == 0 {
			httpx.Fail(c, ErrUnauthenticatedRequest)
			return
		}
		refreshTokens, err := contract.AskRefreshTokensList(c.Request.Context(), resolve(), contract.RefreshTokensList{
			IdentityID: a.IdentityID,
		})
		if err != nil {
			httpx.Fail(c, err)
			return
		}

		refreshTokensDTOs := make([]refreshTokenDTO, 0, len(refreshTokens))
		for _, t := range refreshTokens {
			refreshTokensDTOs = append(refreshTokensDTOs, toRefreshTokenDTO(t))
		}

		httpx.Success(c, refreshTokensDTOs)
	}
}

func handleRefreshTokenDelete(resolve core.Resolve) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			httpx.Fail(c, err)
			return
		}

		ok, err := contract.ExecRefreshTokenDelete(
			c.Request.Context(), resolve(), contract.RefreshTokenDelete{
				TokenID: tokenId,
			},
		)
		if err != nil {
			httpx.Fail(c, err)
			return
		}
		if !ok {
			httpx.Fail(c, errors.New("was not able to delete refresh token"))
			return
		}

		httpx.Success(c, ok)
	}
}

func handleRefreshTokenRevoke(resolve core.Resolve) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			httpx.Fail(c, err)
			return
		}

		token, err := contract.ExecRefreshTokenRevoke(
			c.Request.Context(), resolve(), contract.RefreshTokenRevoke{
				TokenID: tokenId,
			},
		)
		if err != nil {
			httpx.Fail(c, err)
			return
		}

		httpx.Success(c, toRefreshTokenDTO(token))
	}
}
