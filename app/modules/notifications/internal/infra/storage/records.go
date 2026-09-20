package storage

import (
	"time"
)

type notificationRecord struct {
	ID         int        `gorm:"primaryKey;autoIncrement"`
	IdentityID *int       `gorm:"column:identity_id"`
	Type       string     `gorm:"column:type;not null"`
	Content    string     `gorm:"column:content;not null"`
	RefID      *int       `gorm:"column:ref_id"`
	Priority   *string    `gorm:"column:priority"`
	ReadAt     *time.Time `gorm:"column:read_at;default:null"`
	CreatedAt  time.Time  `gorm:"autoCreateTime"`
}

func (notificationRecord) TableName() string { return "notifications" }
