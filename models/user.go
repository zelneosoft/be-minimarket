package models

type User struct {
	ID       uint   `gorm:"column:id;primaryKey" json:"id,omitempty"`
	Email    string `gorm:"column:email" json:"email,omitempty"`
	Password string `gorm:"column:password" json:"-"`
	Name     string `gorm:"column:name" json:"name,omitempty"`
	Level    int    `gorm:"column:level" json:"level,omitempty"`
	IsActive int    `gorm:"column:is_active" json:"is_active,omitempty"`
}
