package contract

import "time"

type Registered struct {
	IdentityID int
	Email      string
	At         time.Time
}

func (Registered) EventName() string {
	return "identity.registered"
}
