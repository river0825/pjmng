package table_del

import (
	"time"

	"gorm.io/datatypes"
)

type Material struct {
	ID          int            `gorm:"column:id;primaryKey;autoIncrement:true"`
	Name        string         `gorm:"column:name"`
	Tag         string         `gorm:"column:tag"`
	EditAt      *time.Time     `gorm:"column:editAt"`
	User        *int           `gorm:"column:user;index:materials_user_idx"`
	PublishedAt *time.Time     `gorm:"column:published_at"`
	CreatedBy   *int           `gorm:"column:created_by"`
	UpdatedBy   *int           `gorm:"column:updated_by"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
	Element     *int           `gorm:"column:element"`
	Content     datatypes.JSON `gorm:"column:content"`
}

func (Material) TableName() string {
	return "materials"
}
