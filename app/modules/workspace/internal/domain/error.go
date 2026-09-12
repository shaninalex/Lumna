package domain

import "gitlab.com/shaninalex/lumna/app/core/errs"

var (
	ErrWorkspaceNotFound = errs.NotFound("WCS001", "workspace not found")
)
