// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package table

const TableNameEditorImagesCategorie = "editor_images__categories"

// EditorImagesCategorie mapped from table <editor_images__categories>
type EditorImagesCategorie struct {
	ID                    int32  `gorm:"column:id;type:integer;not null" json:"id"`
	EditorImageID         *int32 `gorm:"column:editor_image_id;type:integer" json:"editor_image_id"`
	EditorCategoryImageID *int32 `gorm:"column:editor-category-image_id;type:integer" json:"editor-category-image_id"`
}

// TableName EditorImagesCategorie's table name
func (*EditorImagesCategorie) TableName() string {
	return TableNameEditorImagesCategorie
}
