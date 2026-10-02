package activity

import (
	"time"

	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
)

type activityDTO struct {
	Id         int       `json:"id"`
	Message    string    `json:"message"`
	EventType  string    `json:"event_type"`
	EntityId   int       `json:"entity_id"`
	EntityType string    `json:"entity_type"`
	CreatedAt  time.Time `json:"created_at"`
}

func toActivityDTO(s contract.ActivityView) activityDTO {
	return activityDTO{
		Id:         s.Id,
		Message:    s.Message,
		EventType:  s.EventType,
		EntityId:   s.EntityId,
		EntityType: s.EntityType,
		CreatedAt:  s.CreatedAt,
	}
}
