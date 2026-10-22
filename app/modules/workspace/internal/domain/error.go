package domain

import "gitlab.com/shaninalex/lumna/app/core/errs"

var (
	ErrWorkspaceNotFound                 = errs.NotFound("WSP00", "workspace not found")
	CreateInvitationAlreadyExistsError   = errs.Conflict("WSP01", "user already invited in this workspace")
	CreateInvitationAlreadyAcceptedError = errs.Conflict("WSP02", "user already accepted in this workspace")
	CreateInvitationAlreadyRevokedError  = errs.Conflict("WSP03", "invitation was revoked for this user")
	InvitationNotFoundError              = errs.NotFound("WSP04", "invitation not found")
	InvitationExpiredOrAcceptedError     = errs.Conflict("WSP05", "invitation expired or already accepted")
)
