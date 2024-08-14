// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package table

import (
	"time"

	"gorm.io/datatypes"
)

const TableNameShopifyEvent = "shopify_events"

// ShopifyEvent mapped from table <shopify_events>
type ShopifyEvent struct {
	ID        int32           `gorm:"column:id;type:integer;primaryKey;autoIncrement:true" json:"id"`
	EventID   string          `gorm:"column:event_id;type:character varying(255);not null" json:"event_id"`
	Status    *string         `gorm:"column:status;type:character varying(255)" json:"status"`
	EventBody *datatypes.JSON `gorm:"column:event_body;type:jsonb;not null" json:"event_body"`
	CreatedBy *int32          `gorm:"column:created_by;type:integer" json:"created_by"`
	UpdatedBy *int32          `gorm:"column:updated_by;type:integer" json:"updated_by"`
	CreatedAt *time.Time      `gorm:"column:created_at;type:timestamp with time zone;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time      `gorm:"column:updated_at;type:timestamp with time zone;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName ShopifyEvent's table name
func (*ShopifyEvent) TableName() string {
	return TableNameShopifyEvent
}
