package table_del

// DefaultFreeMaterialsAdminMaterials represents the junction table
type DefaultFreeMaterialsAdminMaterials struct {
	ID                    int  `gorm:"column:id;primaryKey;autoIncrement:true"`
	DefaultFreeMaterialID *int `gorm:"column:default_free_material_id"`
	AdminMaterialID       *int `gorm:"column:admin-material_id"`
}

// TableName specifies the table name for DefaultFreeMaterialsAdminMaterials
func (DefaultFreeMaterialsAdminMaterials) TableName() string {
	return "default_free_materials__admin_materials"
}
