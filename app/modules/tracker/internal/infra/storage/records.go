package storage

import (
	"database/sql"
	"time"

	"gitlab.com/shaninalex/lumna/app/modules/tracker/internal/domain"
)

type stageRecord struct {
	ID          int `gorm:"primaryKey;autoIncrement"`
	ScopeID     int
	Name        string
	Description sql.NullString
	Category    string
	Position    float64
	WipLimit    sql.NullInt32 `gorm:"column:wip_limit"`

	Scope     scopeRecord      `gorm:"foreignKey:ScopeID;references:ID"`
	WorkItems []workItemRecord `gorm:"foreignKey:StageID;references:ID"`

	CreatedAt time.Time    `gorm:"autoCreateTime"`
	UpdatedAt sql.NullTime `gorm:"column:updated_at;default:null"`
}

func (stageRecord) TableName() string { return "stages" }

func stageRecordToDomain(record stageRecord) domain.Stage {
	st := domain.Stage{
		ID:          record.ID,
		ScopeID:     record.ScopeID,
		Name:        record.Name,
		Description: record.Description.String,
		Category:    domain.StageCategory(record.Category),
		Position:    record.Position,
		WIPLimit:    int(record.WipLimit.Int32),
		CreatedAt:   record.CreatedAt,
		WorkItems:   workItemsToDomain(record.WorkItems),
	}

	if record.UpdatedAt.Valid {
		st.UpdatedAt = record.UpdatedAt.Time
	}

	return st
}

type scopeRecord struct {
	ID          int `gorm:"primaryKey;autoIncrement"`
	ProjectID   int
	Name        string
	Description sql.NullString

	Stages []stageRecord `gorm:"foreignKey:ScopeID;references:ID"`

	CreatedAt time.Time    `gorm:"autoCreateTime"`
	UpdatedAt sql.NullTime `gorm:"column:updated_at;default:null"`
}

func (scopeRecord) TableName() string { return "scopes" }

func scopeRecordToDomain(record scopeRecord) domain.Scope {
	stages := make([]domain.Stage, len(record.Stages))
	for i, stage := range record.Stages {
		stages[i] = stageRecordToDomain(stage)
	}
	sc := domain.Scope{
		ID:          record.ID,
		ProjectID:   record.ProjectID,
		Name:        record.Name,
		Description: record.Description.String,
		Stages:      stages,
		CreatedAt:   record.CreatedAt,
	}
	if record.UpdatedAt.Valid {
		sc.UpdatedAt = record.UpdatedAt.Time
	}
	return sc
}

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
	ParentID    sql.NullInt64
	Title       string
	Description sql.NullString
	ScopeID     sql.NullInt64
	StageID     sql.NullInt64
	Rank        float64

	//	Priority    *string
	//	SprintID    *int
	//	Estimate    *string
	Assignees []workItemAssignRecord `gorm:"foreignKey:WorkItemID;references:ID"`
	//	Parent   *workItemRecord  `gorm:"foreignKey:ParentID;references:ID"`
	//	Children []workItemRecord `gorm:"foreignKey:ParentID;references:ID"`

	DueTo     sql.NullTime `gorm:"column:due_to;default:null"`
	CreatedAt time.Time    `gorm:"autoCreateTime"`
	UpdatedAt sql.NullTime `gorm:"column:updated_at;default:null"`
}

func (workItemRecord) TableName() string { return "work_items" }

func workItemToDomain(record workItemRecord) domain.WorkItem {
	assigneeIDs := make([]int, len(record.Assignees))
	for j, assignee := range record.Assignees {
		assigneeIDs[j] = assignee.IdentityID
	}
	wt := domain.WorkItem{
		ID:          record.ID,
		ProjectID:   record.ProjectID,
		Type:        domain.WorkItemType(record.Type),
		ParentID:    int(record.ParentID.Int64),
		Title:       record.Title,
		Description: record.Description.String,
		ScopeID:     int(record.ScopeID.Int64),
		StageID:     int(record.StageID.Int64),
		Rank:        record.Rank,
		CreatedAt:   record.CreatedAt,
		AssigneeIDs: assigneeIDs,
	}

	if record.UpdatedAt.Valid {
		wt.UpdatedAt = record.UpdatedAt.Time
	}

	if record.DueTo.Valid {
		wt.UpdatedAt = record.DueTo.Time
	}

	return wt
}

func workItemsToDomain(records []workItemRecord) []domain.WorkItem {
	workItems := make([]domain.WorkItem, len(records))
	for i, record := range records {
		workItems[i] = workItemToDomain(record)
	}
	return workItems
}
