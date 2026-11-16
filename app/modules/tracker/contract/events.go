package contract

import journal "gitlab.com/shaninalex/lumna/app/modules/journal/contract"

const entityWorkItem = "work_item"

var (
	_ journal.Recordable = WorkItemStageChanged{}
	_ journal.Recordable = WorkItemAssigned{}
	_ journal.Recordable = WorkItemUnassigned{}
)

type WorkItemStageChanged struct {
	WorkItemId int
	StageId    int
	ActorId    int
	// AssigneeIds are the work item's assignees at the moment of the change.
	AssigneeIds []int
}

func (e WorkItemStageChanged) EventName() string             { return "work_item/stage_changed" }
func (e WorkItemStageChanged) JournalSubject() (string, int) { return entityWorkItem, e.WorkItemId }
func (e WorkItemStageChanged) JournalActor() int             { return e.ActorId }

type WorkItemAssigned struct {
	WorkItemId int
	AssigneeId int
	ActorId    int
}

func (e WorkItemAssigned) EventName() string             { return "work_item/assigned" }
func (e WorkItemAssigned) JournalSubject() (string, int) { return entityWorkItem, e.WorkItemId }
func (e WorkItemAssigned) JournalActor() int             { return e.ActorId }

type WorkItemUnassigned struct {
	WorkItemId int
	AssigneeId int
	ActorId    int
}

func (e WorkItemUnassigned) EventName() string             { return "work_item/unassigned" }
func (e WorkItemUnassigned) JournalSubject() (string, int) { return entityWorkItem, e.WorkItemId }
func (e WorkItemUnassigned) JournalActor() int             { return e.ActorId }
