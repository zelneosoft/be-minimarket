package models

type Category struct {
	ID       uint   `gorm:"column:id;primaryKey" json:"id,omitempty"`
	Name     string `gorm:"column:name" json:"name,omitempty"`
	IsActive int    `gorm:"column:is_active" json:"is_active,omitempty"`
	Color    string `gorm:"column:color" json:"color,omitempty"`
}
