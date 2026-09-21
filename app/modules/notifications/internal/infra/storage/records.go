package storage

import (
	"database/sql"
	"time"
)

type notificationRecord struct {
	ID         int            `gorm:"primaryKey;autoIncrement"`
	IdentityID sql.NullInt64  `gorm:"column:identity_id"`
	Type       string         `gorm:"column:type;not null"`
	Content    string         `gorm:"column:content;not null"`
	RefID      sql.NullInt64  `gorm:"column:ref_id"`
	Priority   sql.NullString `gorm:"column:priority"`
	ReadAt     sql.NullTime   `gorm:"column:read_at;default:null"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
}

func (notificationRecord) TableName() string { return "notifications" }
