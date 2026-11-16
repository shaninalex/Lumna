package domain

import (
	"slices"
	"testing"
)

func TestNotification_AddRecipients(t *testing.T) {
	var n Notification
	n.AddRecipients(7, 1, 7, 0, 2, 1)
	n.AddRecipients(7, 2, 3)

	if got, want := n.RecipientIDs(), []int{1, 2, 3}; !slices.Equal(got, want) {
		t.Fatalf("recipients = %v, want %v", got, want)
	}
}
