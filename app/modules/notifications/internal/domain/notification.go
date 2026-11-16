package domain

import (
	"slices"
	"time"
)

type Notification struct {
	ID               int
	NotificationType string
	// Content is JSON data for the client to render the message from, so the
	// text is not frozen in one language at write time.
	Content    string
	RefId      int
	RefType    string
	Priority   string
	Created    time.Time
	Recipients []Recipient
}

type Recipient struct {
	ID             int
	NotificationID int
	IdentityID     int
	ReadAt         time.Time
}

// AddRecipients adds identities as recipients, skipping zero ids, the actor
// who caused the notification and identities that are already recipients.
func (n *Notification) AddRecipients(actorID int, identityIDs ...int) {
	for _, id := range identityIDs {
		if id == 0 || id == actorID || slices.Contains(n.RecipientIDs(), id) {
			continue
		}
		n.Recipients = append(n.Recipients, Recipient{IdentityID: id})
	}
}

func (n Notification) RecipientIDs() []int {
	ids := make([]int, len(n.Recipients))
	for i, r := range n.Recipients {
		ids[i] = r.IdentityID
	}
	return ids
}
