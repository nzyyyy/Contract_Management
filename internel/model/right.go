package model

type Right struct {
	UserId      int `gorm:"column:user_id"`
	RoleId      int `gorm:"column:role_id"`
	Description string
}

func (r Right) TableName() string {
	return "right"
}
