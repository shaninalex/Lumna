package domain

import "gitlab.com/shaninalex/lumna/app/core/errs"

var (
	ErrBadCredentials      = errs.Unauthenticated("AUTH001", "invalid credentials")
	ErrTokenUnknown        = errs.Unauthenticated("AUTH002", "refresh token is not recognized")
	ErrTokenExpired        = errs.Unauthenticated("AUTH003", "refresh token has expired")
	ErrTokenRevoked        = errs.Unauthenticated("AUTH004", "refresh token has been revoked")
	ErrEmptyToken          = errs.Validation("AUTH005", "refresh token is empty")
	ErrBadAccessToken      = errs.Unauthenticated("AUTH006", "access token is missing or invalid")
	ErrAccessClaimsInvalid = errs.Validation("AUTH007", "invalid access claims")
	ErrTokenNotBelongs     = errs.Conflict("AUTH008", "token not belongs to given actor")
	ErrNotFound            = errs.Conflict("AUTH009", "not found")
)
