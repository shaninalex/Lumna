package column

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type columnDto struct {
	Id        int        `json:"id"`
	Title     string     `json:"title"`
	Meta      ColumnMeta `json:"meta"`
	BoardId   int        `json:"board_id"`
	Position  float64    `json:"position"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type ColumnMeta struct {
	Color    string `json:"color"`
	Icon     string `json:"icon"`
	Expanded bool   `json:"expanded"`
}

func newDefaultColumnMeta() ColumnMeta {
	return ColumnMeta{
		Color:    "default",
		Icon:     "default",
		Expanded: true,
	}
}

func (j *ColumnMeta) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("failed to unmarshal JSONB value:", value))
	}

	result := ColumnMeta{}
	err := json.Unmarshal(bytes, &result)
	*j = result
	return err
}

func (j ColumnMeta) Value() (driver.Value, error) {
	b, err := json.Marshal(j)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(b).MarshalJSON()
}

type columnCreateDto struct {
	Title    string  `json:"title"`
	Position float64 `json:"order"`
	BoardId  int     `json:"board_id"`
}

func (a columnCreateDto) Validate() error {
	return validation.ValidateStruct(
		&a,
		validation.Field(&a.Title, validation.Required, validation.Length(3, 50)),
		validation.Field(&a.BoardId, validation.Required, validation.Min(0)),
	)
}

type columnDeleteResponseDto struct {
	Id           int   `json:"id"`
	DeletedTasks []int `json:"deleted_tasks"`
}
