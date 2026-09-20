package contract

type WorkItemStageChanged struct {
	WorkItemId int
	StageId    int
}

func (e WorkItemStageChanged) EventName() string {
	return "work_item/stage_changed"
}
