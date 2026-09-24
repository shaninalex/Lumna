package contract

import "time"

type NotificationView struct {
	ID        int
	Type      string
	Content   string
	RefID     int
	CreatedAt time.Time
}
