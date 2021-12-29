package model

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	RoleId   int    `gorm:"column:roleId"`
}

func (u User) TableName() string {
	return "user"
}
