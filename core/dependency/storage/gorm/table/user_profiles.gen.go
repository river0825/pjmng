// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package table

import (
	"time"
)

const TableNameUserProfile = "user_profiles"

// UserProfile mapped from table <user_profiles>
type UserProfile struct {
	ID          int32      `gorm:"column:id;type:integer;primaryKey;autoIncrement:true" json:"id"`
	Institute   *string    `gorm:"column:institute;type:character varying(255)" json:"institute"`
	Tenure      *string    `gorm:"column:tenure;type:character varying(255)" json:"tenure"`
	User        *int32     `gorm:"column:user;type:integer" json:"user"`
	Name        *string    `gorm:"column:name;type:character varying(255)" json:"name"`
	Role        *string    `gorm:"column:role;type:character varying(255)" json:"role"`
	PublishedAt *time.Time `gorm:"column:published_at;type:timestamp with time zone" json:"published_at"`
	CreatedBy   *int32     `gorm:"column:created_by;type:integer" json:"created_by"`
	UpdatedBy   *int32     `gorm:"column:updated_by;type:integer" json:"updated_by"`
	CreatedAt   *time.Time `gorm:"column:created_at;type:timestamp with time zone;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at;type:timestamp with time zone;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName UserProfile's table name
func (*UserProfile) TableName() string {
	return TableNameUserProfile
}