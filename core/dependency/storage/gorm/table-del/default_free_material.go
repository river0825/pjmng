package table_del

import "time"

// DefaultFreeMaterial represents the default_free_materials table
type DefaultFreeMaterial struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement:true"`
	CreatedBy *int      `gorm:"column:created_by"`
	UpdatedBy *int      `gorm:"column:updated_by"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	// Many-to-Many relationship
	AdminMaterials []AdminMaterial `gorm:"many2many:default_free_materials__admin_materials;foreignKey:ID;joinForeignKey:default_free_material_id;References:ID;joinReferences:admin-material_id"`
}

// TableName specifies the table name for DefaultFreeMaterial
func (DefaultFreeMaterial) TableName() string {
	return "default_free_materials"
}
