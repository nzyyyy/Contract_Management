package model

type Role struct {
	ID          int `gorm:"autoIncrement"`
	Name        string
	Description string
	Functions   string
}

func (r Role) TableName() string {
	return "role"
}
