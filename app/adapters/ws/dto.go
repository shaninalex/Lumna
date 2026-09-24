package ws

import (
	"time"

	"gitlab.com/shaninalex/lumna/app/modules/notifications/contract"
	"gitlab.com/shaninalex/lumna/app/platform/realtime"
)

type MessageDTO struct {
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}

type NotificationDTO struct {
	ID        int       `json:"id"`
	Type      string    `json:"type"`
	Content   string    `json:"content"`
	RefID     int       `json:"ref_id"`
	CreatedAt time.Time `json:"created_at"`
}

func encode(e realtime.Envelope) (MessageDTO, bool) {
	switch v := e.Payload.(type) {
	case contract.NotificationView:
		return MessageDTO{Type: e.Name, Payload: toNotificationDTO(v)}, true
	default:
		return MessageDTO{}, false
	}
}

func toNotificationDTO(n contract.NotificationView) NotificationDTO {
	return NotificationDTO{
		ID:        n.ID,
		Type:      n.Type,
		Content:   n.Content,
		RefID:     n.RefID,
		CreatedAt: n.CreatedAt,
	}
}
