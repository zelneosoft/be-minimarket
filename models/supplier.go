package models

type Supplier struct {
	ID       uint   `gorm:"column:id;primaryKey" json:"id,omitempty"`
	Name     string `gorm:"column:name" json:"name,omitempty"`
	Address  string `gorm:"column:address" json:"address,omitempty"`
	Maps     string `gorm:"column:maps" json:"maps,omitempty"`
	Phone    string `gorm:"column:phone" json:"phone,omitempty"`
	IsActive int    `gorm:"column:is_active" json:"is_active,omitempty"`
}
