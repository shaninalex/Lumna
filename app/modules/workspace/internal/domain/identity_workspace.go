package domain

import "time"

type IdentityWorkspace struct {
	IdentityID  int
	WorkspaceID int
	CreatedAt   time.Time
}
