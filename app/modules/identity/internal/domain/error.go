package domain

import "gitlab.com/shaninalex/lumna/app/core/errs"

var (
	ErrInvalidEmail   = errs.Validation("IDN001", "invalid email")
	ErrEmailTaken     = errs.Conflict("IDN002", "email already registered")
	ErrNotFound       = errs.Unauthenticated("IDN003", "not found")
	ErrNotActive      = errs.Unauthenticated("IDN004", "account is not active")
	ErrBadCredentials = errs.Unauthenticated("IDN005", "invalid credentials")
	ErrEmptyName      = errs.Forbidden("IDN006", "empty name")
)
