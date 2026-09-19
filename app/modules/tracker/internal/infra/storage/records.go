package storage

import (
	"time"

	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
)

type stageRecord struct {
	ID          int `gorm:"primaryKey;autoIncrement"`
	ScopeID     int
	Name        string
	Description *string
	Category    string
	Position    float64
	WipLimit    *int `gorm:"column:wip_limit"`

	Scope scopeRecord `gorm:"foreignKey:ScopeID;references:ID"`
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

	Stages []stageRecord `gorm:"foreignKey:ScopeID;references:ID"`

	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
}

func (scopeRecord) TableName() string { return "scopes" }

type workItemAssignRecord struct {
	IdentityID int `gorm:"primaryKey"`
	WorkItemID int `gorm:"primaryKey"`
	//WorkItem workItemRecord `gorm:"foreignKey:WorkItemID;references:ID"`
}

func (workItemAssignRecord) TableName() string { return "work_items_assignees" }

type workItemRecord struct {
	ID          int `gorm:"primaryKey;autoIncrement"`
	ProjectID   int
	Type        string
	ParentID    *int
	Title       string
	Description *string
	ScopeID     *int
	StageID     *int
	Rank        float64

	//	Priority    *string
	//	SprintID    *int
	//	Estimate    *string
	Assignees []workItemAssignRecord `gorm:"foreignKey:WorkItemID;references:ID"`
	//	Parent   *workItemRecord  `gorm:"foreignKey:ParentID;references:ID"`
	//	Children []workItemRecord `gorm:"foreignKey:ParentID;references:ID"`

	DueTo     *time.Time `gorm:"due_to"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
}

func (workItemRecord) TableName() string { return "work_items" }

func workItemToDomain(record workItemRecord) domain.WorkItem {
	assigneeIDs := make([]int, len(record.Assignees))
	for j, assignee := range record.Assignees {
		assigneeIDs[j] = assignee.IdentityID
	}
	return domain.WorkItem{
		ID:          record.ID,
		ProjectID:   record.ProjectID,
		Type:        domain.WorkItemType(record.Type),
		ParentID:    record.ParentID,
		Title:       record.Title,
		Description: record.Description,
		ScopeID:     record.ScopeID,
		StageID:     record.StageID,
		Rank:        record.Rank,
		DueTo:       record.DueTo,
		CreatedAt:   record.CreatedAt,
		UpdatedAt:   record.UpdatedAt,
		AssigneeIDs: assigneeIDs,
	}
}
