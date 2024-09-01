package products

type ProductRequest struct {
	ID          uint   `gorm:"column:id;primaryKey" json:"id,omitempty"`
	Barcode     string `gorm:"column:barcode" json:"barcode,omitempty"`
	Name        string `gorm:"column:name" json:"name,omitempty"`
	Description string `gorm:"column:description" json:"description,omitempty"`
	CategoryID  uint   `gorm:"column:category_id" json:"category_id,omitempty"`
	BrandID     uint   `gorm:"column:brand_id" json:"brand_id,omitempty"`
}
