package user

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/adapters/httpx"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/actor"
	"gitlab.com/shaninalex/lumna/app/core/errs"
	"gitlab.com/shaninalex/lumna/app/core/secure"
	authContract "gitlab.com/shaninalex/lumna/app/modules/auth/contract"
	"gitlab.com/shaninalex/lumna/app/modules/identity/contract"
)

var (
	ErrAuthenticationRequired = errs.Unauthenticated("AUTH001", "authentication required")
	ErrUnableToChangePassword = errs.Conflict("AUTH002", "unable to change password")
	ErrPasswordsDoesNotMatch  = errs.Conflict("AUTH003", "passwords does not match")
	ErrPasswordToWeak         = errs.Conflict("AUTH004", "password too weak")
)

func Register(resolve core.Resolve, router *gin.RouterGroup, verifier authContract.PasswordVerifier) {
	router.GET("/me", handleMe(resolve))
	router.POST("/change-password", handleChangePassword(resolve, verifier))
}

func handleMe(resolve core.Resolve) gin.HandlerFunc {
	return func(c *gin.Context) {
		a, ok := actor.From(c.Request.Context())
		if !ok {
			httpx.Fail(c, ErrAuthenticationRequired)
			return
		}
		identity, err := contract.AskGetProfile(c.Request.Context(), resolve(), contract.GetProfile{IdentityID: a.IdentityID})
		if err != nil {
			httpx.Fail(c, err)
			return
		}

		httpx.Success(c, profileDTO{
			ID:       identity.ID,
			Email:    identity.Email,
			FullName: identity.FullName,
			Active:   identity.Active,
		})
	}
}

func handleChangePassword(resolve core.Resolve, verifier authContract.PasswordVerifier) gin.HandlerFunc {
	return func(c *gin.Context) {
		a, ok := actor.From(c.Request.Context())
		if !ok {
			httpx.Fail(c, errs.Unauthenticated("API001", "authentication required"))
			return
		}
		var data changePasswordDTO
		if err := c.BindJSON(&data); err != nil {
			httpx.Fail(c, err)
			return
		}

		if data.Password != data.PasswordConfirm {
			httpx.Fail(c, ErrPasswordsDoesNotMatch)
			return
		}
		securePassword := secure.Secret(data.Password)
		if verifier.Verify(securePassword) < 4 {
			httpx.Fail(c, ErrPasswordToWeak)
			return
		}

		ok, err := contract.ExecPasswordChange(c.Request.Context(), resolve(), contract.PasswordChange{
			IdentityId: a.IdentityID,
			Password:   securePassword,
		})
		if err != nil {
			httpx.Fail(c, err)
			return
		}
		if !ok {
			httpx.Fail(c, ErrUnableToChangePassword)
			return
		}

		httpx.Success(c, true)
	}
}
