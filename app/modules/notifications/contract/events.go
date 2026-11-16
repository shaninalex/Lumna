package contract

// NotificationCreated is published after a notification and its recipients
// are stored.
type NotificationCreated struct {
	NotificationId int
	Type           string
	RecipientIds   []int
}

func (n NotificationCreated) EventName() string {
	return "notification/created"
}
