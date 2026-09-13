package storage

import "time"

type stageRecord struct {
	ID          int `gorm:"primaryKey;autoIncrement"`
	ScopeID     int `gorm:"foreignKey:scope_id"`
	Name        string
	Description *string
	Category    string
	Position    float64
	WipLimit    *int `gorm:"column:wip_limit"`

	Scope scopeRecord
	//WorkItems []workItemRecord

	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
}

func (stageRecord) TableName() string { return "stages" }

type scopeRecord struct {
	ID          int `gorm:"primaryKey;autoIncrement"`
	ProjectID   int
	Name        string
	Description *string

	Stages []stageRecord `gorm:"many2many:stages"`

	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
}

func (scopeRecord) TableName() string { return "scopes" }

//type sprintRecord struct {
//	ID          int `gorm:"primaryKey;autoIncrement"`
//	ProjectID   int
//	Name        string
//	Description *string
//	StartDate   *time.Time
//	EndDate     *time.Time
//	IsActive    bool
//
//	WorkItems []workItemRecord
//
//	CreatedAt time.Time  `gorm:"autoCreateTime"`
//	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
//}
//
//type workItemRecord struct {
//	ID          int `gorm:"primaryKey;autoIncrement"`
//	ProjectID   int
//	Type        string
//	ParentID    *int
//	Title       string
//	Description *string
//	ScopeID     *int
//	StageID     *int
//	Rank        *float64
//	Priority    *string
//	SprintID    *int
//	Estimate    *string
//
//	Parent   *workItemRecord  `gorm:"foreignKey:ParentID;references:ID"`
//	Children []workItemRecord `gorm:"foreignKey:ParentID;references:ID"`
//
//	Scope  *scopeRecord
//	Stage  *stageRecord
//	Sprint *sprintRecord
//
//	Assignees []workItemAssigneeRecord `gorm:"foreignKey:WorkItemID;references:ID"`
//
//	CreatedAt time.Time  `gorm:"autoCreateTime"`
//	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
//}
//
//type workItemAssigneeRecord struct {
//	WorkItemID int `gorm:"primaryKey"`
//	IdentityID int `gorm:"primaryKey"`
//}
//
//type workItemRelationRecord struct {
//	SourceID int    `gorm:"primaryKey"`
//	TargetID int    `gorm:"primaryKey"`
//	Type     string `gorm:"primaryKey"`
//
//	Source workItemRecord `gorm:"foreignKey:SourceID;references:ID"`
//	Target workItemRecord `gorm:"foreignKey:TargetID;references:ID"`
//
//	CreatedAt time.Time `gorm:"autoCreateTime"`
//}
