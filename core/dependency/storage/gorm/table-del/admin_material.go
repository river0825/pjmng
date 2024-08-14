package table_del

import (
	"time"

	"gorm.io/datatypes"
)

type AdminMaterial struct {
	ID           int            `gorm:"column:id;primaryKey;autoIncrement:true"`
	Name         string         `gorm:"column:name;type:varchar(255)"`
	Tag          string         `gorm:"column:tag;type:varchar(255)"`
	EditAt       *time.Time     `gorm:"column:editAt;type:timestamp"`
	User         *int           `gorm:"column:user"`
	Content      datatypes.JSON `gorm:"column:content;type:jsonb"`
	Lesson       *int           `gorm:"column:lesson"`
	Grammar      *int           `gorm:"column:grammer"`
	PublishedAt  *time.Time     `gorm:"column:published_at;type:timestamp"`
	CreatedBy    *int           `gorm:"column:created_by"`
	UpdatedBy    *int           `gorm:"column:updated_by"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	Organization *int           `gorm:"column:organization"`
	TagSimplify  string         `gorm:"column:tag_simplify;type:varchar(255)"`
	IsFree       bool           `gorm:"column:isfree"`

	// Many-to-Many relationship
	DefaultFreeMaterials []DefaultFreeMaterial `gorm:"many2many:default_free_materials__admin_materials;foreignKey:ID;joinForeignKey:admin-material_id;References:ID;joinReferences:default_free_material_id"`
}

func (AdminMaterial) TableName() string {
	return "admin_materials"
}
