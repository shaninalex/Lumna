package dto

import (
	"time"

	"gitlab.com/shaninalex/lumna/app/models"
)

type EntityEventDTO struct {
	ID         int       `json:"id"`
	IdentityId *int      `json:"identity_id,omitempty"`
	EntityId   *int      `json:"entity_id,omitempty"`
	EntityType *string   `json:"entity_type,omitempty"`
	EventType  string    `json:"event_type"`
	Data       string    `json:"data"`
	CreatedAt  time.Time `json:"created_at"`
}

func ToEntityEventDTO(e models.EntityEvent) EntityEventDTO {
	return EntityEventDTO{
		ID:         e.ID,
		IdentityId: e.IdentityId,
		EntityId:   e.EntityId,
		EntityType: e.EntityType,
		EventType:  string(e.EventType),
		Data:       e.Data,
		CreatedAt:  e.CreatedAt,
	}
}
