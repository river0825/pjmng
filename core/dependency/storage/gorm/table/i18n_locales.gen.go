// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package table

import (
	"time"
)

const TableNameI18nLocale = "i18n_locales"

// I18nLocale mapped from table <i18n_locales>
type I18nLocale struct {
	ID        int32      `gorm:"column:id;type:integer;not null" json:"id"`
	Name      *string    `gorm:"column:name;type:character varying(255)" json:"name"`
	Code      *string    `gorm:"column:code;type:character varying(255)" json:"code"`
	CreatedBy *int32     `gorm:"column:created_by;type:integer" json:"created_by"`
	UpdatedBy *int32     `gorm:"column:updated_by;type:integer" json:"updated_by"`
	CreatedAt *time.Time `gorm:"column:created_at;type:timestamp with time zone;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:timestamp with time zone;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName I18nLocale's table name
func (*I18nLocale) TableName() string {
	return TableNameI18nLocale
}
