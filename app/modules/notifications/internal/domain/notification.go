package domain

import "time"

type Notification struct {
	ID               int
	IdentityId       *int
	NotificationType string
	Content          string
	RefId            *int
	Priority         *string
	ReadAt           *time.Time
	Created          time.Time
}
