package dto

import (
	"time"

	"gitlab.com/shaninalex/lumna/app/models"
)

type ColumnDTO struct {
	ID        int               `json:"id"`
	Title     string            `json:"title"`
	Meta      models.ColumnMeta `json:"meta"`
	BoardId   int               `json:"board_id"`
	Position  int               `json:"position"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt *time.Time        `json:"updated_at"`
}

func ColumnDTOToModel(c ColumnDTO) models.Column {
	return models.Column{
		ID:        c.ID,
		Title:     c.Title,
		Meta:      c.Meta,
		BoardId:   c.BoardId,
		Position:  c.Position,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

type ColumnReorderingDto struct {
	ColumnID      int   `json:"id"`
	PreviousIndex int   `json:"previous_index"`
	CurrentIndex  int   `json:"current_index"`
	BoardId       int   `json:"board_id"`
	ColumnsOrder  []int `json:"columns_order"`
}
