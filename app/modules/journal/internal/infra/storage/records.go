package storage

import (
	"database/sql"
	"time"
)

type entryRecord struct {
	ID         int            `gorm:"primaryKey;autoIncrement"`
	ActorID    sql.NullInt64  `gorm:"column:actor_id"`
	EntityType string         `gorm:"column:entity_type;not null"`
	EntityID   int            `gorm:"column:entity_id;not null"`
	EventType  string         `gorm:"column:event_type;not null"`
	Data       sql.NullString `gorm:"column:data"`
	CreatedAt  time.Time      `gorm:"column:created_at"`
}

func (entryRecord) TableName() string { return "journal" }
