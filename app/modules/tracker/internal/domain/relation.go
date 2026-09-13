package domain

type RelationType string

const (
	RelationBlocks    RelationType = "blocks"
	RelationBlockedBy RelationType = "blocked_by"
	RelationRelatesTo RelationType = "relates_to"
	RelationDuplicate RelationType = "duplicate"
)

type WorkItemRelation struct {
	SourceID int // From
	TargetID int // To
	Type     RelationType
}
