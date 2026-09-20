package contract

type NotificationCreated struct {
	Source string
	// todo: rest of the fields
}

func (n NotificationCreated) EventName() string {
	return "notification/created"
}
