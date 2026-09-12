package contract

import "time"

type WorkspaceView struct {
	Id         int
	Title      string
	OwnerEmail string
	Active     bool
	CreatedAt  time.Time
	UpdatedAt  *time.Time
}
