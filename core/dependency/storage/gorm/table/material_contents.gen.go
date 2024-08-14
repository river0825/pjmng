// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package table

import (
	"time"

	"gorm.io/datatypes"
)

const TableNameMaterialContent = "material_contents"

// MaterialContent mapped from table <material_contents>
type MaterialContent struct {
	ID        int32           `gorm:"column:id;type:integer;primaryKey;autoIncrement:true" json:"id"`
	PageID    string          `gorm:"column:page_id;type:character varying(255);not null" json:"page_id"`
	Content   *datatypes.JSON `gorm:"column:content;type:jsonb" json:"content"`
	Order_    *int32          `gorm:"column:order;type:integer" json:"order"`
	CreatedBy *int32          `gorm:"column:created_by;type:integer" json:"created_by"`
	UpdatedBy *int32          `gorm:"column:updated_by;type:integer" json:"updated_by"`
	CreatedAt *time.Time      `gorm:"column:created_at;type:timestamp with time zone;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time      `gorm:"column:updated_at;type:timestamp with time zone;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName MaterialContent's table name
func (*MaterialContent) TableName() string {
	return TableNameMaterialContent
}
