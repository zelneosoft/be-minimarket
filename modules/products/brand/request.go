package brand

type BrandRequest struct {
	ID          uint   `gorm:"column:id;primaryKey" json:"id,omitempty"`
	Name        string `gorm:"column:name" json:"name,omitempty"`
	Description string `gorm:"column:description" json:"description,omitempty"`
	IsActive    int    `gorm:"column:is_active" json:"is_active,omitempty"`
}
