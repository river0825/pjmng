// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package table

const TableNameComponentsMaterialPage = "components_material_pages"

// ComponentsMaterialPage mapped from table <components_material_pages>
type ComponentsMaterialPage struct {
	ID         int32   `gorm:"column:id;type:integer;not null" json:"id"`
	Priority   *int32  `gorm:"column:priority;type:integer" json:"priority"`
	Annotation *string `gorm:"column:annotation;type:text" json:"annotation"`
	Content    *string `gorm:"column:content;type:text" json:"content"`
}

// TableName ComponentsMaterialPage's table name
func (*ComponentsMaterialPage) TableName() string {
	return TableNameComponentsMaterialPage
}
