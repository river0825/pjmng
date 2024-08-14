// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package table

import (
	"time"
)

const TableNameEditorAudio = "editor_audios"

// EditorAudio mapped from table <editor_audios>
type EditorAudio struct {
	ID          int32      `gorm:"column:id;type:integer;not null" json:"id"`
	Name        *string    `gorm:"column:name;type:character varying(255)" json:"name"`
	Tag         *string    `gorm:"column:tag;type:character varying(255)" json:"tag"`
	PublishedAt *time.Time `gorm:"column:published_at;type:timestamp with time zone" json:"published_at"`
	CreatedBy   *int32     `gorm:"column:created_by;type:integer" json:"created_by"`
	UpdatedBy   *int32     `gorm:"column:updated_by;type:integer" json:"updated_by"`
	CreatedAt   *time.Time `gorm:"column:created_at;type:timestamp with time zone;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at;type:timestamp with time zone;default:CURRENT_TIMESTAMP" json:"updated_at"`
	Category    *int32     `gorm:"column:category;type:integer" json:"category"`
	TagSimplify *string    `gorm:"column:tag_simplify;type:character varying(255)" json:"tag_simplify"`
}

// TableName EditorAudio's table name
func (*EditorAudio) TableName() string {
	return TableNameEditorAudio
}
