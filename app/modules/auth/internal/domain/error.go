package domain

import "gitlab.com/shaninalex/lumna/app/core/errs"

var (
	ErrBadCredentials = errs.Unauthenticated("AUT001", "invalid credentials")
	ErrTokenUnknown   = errs.Unauthenticated("AUT002", "refresh token is not recognized")
	ErrTokenExpired   = errs.Unauthenticated("AUT003", "refresh token has expired")
	ErrTokenRevoked   = errs.Unauthenticated("AUT004", "refresh token has been revoked")
	ErrEmptyToken     = errs.Validation("AUT005", "refresh token is empty")
	ErrBadAccessToken = errs.Unauthenticated("AUT006", "access token is missing or invalid")
)
