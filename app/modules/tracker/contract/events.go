package contract

import journal "gitlab.com/shaninalex/lumna/app/modules/journal/contract"

const entityWorkItem = "work_item"
const entityStage = "stage"

var (
	_ journal.Recordable = WorkItemStageChanged{}
	_ journal.Recordable = WorkItemAssigned{}
	_ journal.Recordable = WorkItemUnassigned{}
	_ journal.Recordable = WorkItemChanged{}
	_ journal.Recordable = WorkItemMoved{}
)

type WorkItemChanged struct {
	WorkItemId int
	ActorId    int
}

func (e WorkItemChanged) EventName() string             { return "work_item/changed" }
func (e WorkItemChanged) JournalSubject() (string, int) { return entityWorkItem, e.WorkItemId }
func (e WorkItemChanged) JournalActor() int             { return e.ActorId }

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

type WorkItemDeleted struct {
	WorkItemId int
	ActorId    int
}

func (e WorkItemDeleted) EventName() string             { return "work_item/deleted" }
func (e WorkItemDeleted) JournalSubject() (string, int) { return entityWorkItem, e.WorkItemId }
func (e WorkItemDeleted) JournalActor() int             { return e.ActorId }

type WorkItemCreated struct {
	WorkItemId int
	ActorId    int
}

func (e WorkItemCreated) EventName() string             { return "work_item/created" }
func (e WorkItemCreated) JournalSubject() (string, int) { return entityWorkItem, e.WorkItemId }
func (e WorkItemCreated) JournalActor() int             { return e.ActorId }

type WorkItemMoved struct {
	WorkItemId int
	Assignees  []int
	ActorId    int
}

func (e WorkItemMoved) EventName() string             { return "work_item/moved" }
func (e WorkItemMoved) JournalSubject() (string, int) { return entityWorkItem, e.WorkItemId }
func (e WorkItemMoved) JournalActor() int             { return e.ActorId }

type StageMoved struct {
	StageID int
	ActorID int
}

func (e StageMoved) EventName() string             { return "stage/moved" }
func (e StageMoved) JournalSubject() (string, int) { return entityStage, e.StageID }
func (e StageMoved) JournalActor() int             { return e.ActorID }
