package domain

import (
	"time"
)

type Notification struct {
	ID int

	// IdentityId represents the target user/recipient.
	// NOTE: But if there are 0 (NULL) it's broadcast notification.
	// So how then recipients will mark them "as read" ???
	IdentityId       int
	NotificationType string
	Content          string
	RefId            int
	Priority         string
	ReadAt           time.Time
	Created          time.Time
}
