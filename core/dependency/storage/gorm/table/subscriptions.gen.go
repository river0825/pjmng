// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package table

import (
	"time"
)

const TableNameSubscription = "subscriptions"

// Subscription mapped from table <subscriptions>
type Subscription struct {
	ID        int32      `gorm:"column:id;type:integer;not null" json:"id"`
	User      *int32     `gorm:"column:user;type:integer" json:"user"`
	Plan      *int32     `gorm:"column:plan;type:integer" json:"plan"`
	StartAt   *time.Time `gorm:"column:start_at;type:timestamp with time zone" json:"start_at"`
	EndAt     *time.Time `gorm:"column:end_at;type:timestamp with time zone" json:"end_at"`
	CreatedBy *int32     `gorm:"column:created_by;type:integer" json:"created_by"`
	UpdatedBy *int32     `gorm:"column:updated_by;type:integer" json:"updated_by"`
	CreatedAt *time.Time `gorm:"column:created_at;type:timestamp with time zone;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:timestamp with time zone;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName Subscription's table name
func (*Subscription) TableName() string {
	return TableNameSubscription
}
