package supplier

type BranchRequest struct {
	ID       uint   `gorm:"column:id;primaryKey" json:"id,omitempty"`
	Name     string `gorm:"column:name" json:"name,omitempty"`
	Address  string `gorm:"column:address" json:"address,omitempty"`
	Maps     string `gorm:"column:maps" json:"maps,omitempty"`
	Email    string `gorm:"column:email" json:"email,omitempty"`
	Phone    string `gorm:"column:phone" json:"phone,omitempty"`
	Pic      string `gorm:"column:pic" json:"pic,omitempty"`
	IsActive int    `gorm:"column:is_active" json:"is_active,omitempty"`
}
