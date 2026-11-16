package storage

import (
	"database/sql"
	"time"
)

type notificationRecord struct {
	ID        int            `gorm:"primaryKey;autoIncrement"`
	Type      string         `gorm:"column:type;not null"`
	Content   string         `gorm:"column:content;not null"`
	RefID     sql.NullInt64  `gorm:"column:ref_id"`
	RefType   sql.NullString `gorm:"column:ref_type"`
	Priority  sql.NullString `gorm:"column:priority"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
}

func (notificationRecord) TableName() string { return "notifications" }

type notificationRecipientRecord struct {
	ID             int          `gorm:"primaryKey;autoIncrement"`
	IdentityID     int          `gorm:"column:identity_id"`
	NotificationID int          `gorm:"column:notification_id"`
	ReadAt         sql.NullTime `gorm:"column:read_at;default:null"`
}

func (notificationRecipientRecord) TableName() string { return "notification_recipients" }

type recordEmail struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	Subject   string
	Html      string
	Text      string
	Type      string
	Status    string `gorm:"default:pending"`
	Attempts  sql.NullInt16
	Receivers string
	Message   sql.NullString
	SendAt    sql.NullTime `gorm:"column:send_at;default:null"`
	CreatedAt time.Time    `gorm:"autoCreateTime"`
}

func (recordEmail) TableName() string { return "email_queue" }
