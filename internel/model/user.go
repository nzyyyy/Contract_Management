package model

import "gorm.io/gorm"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	RoleId   int    `gorm:"column:roleId"`
}

type APIUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func (u User) TableName() string {
	return "user"
}

func (u User) ListByRole(db *gorm.DB) ([]*APIUser, error) {
	var result []*APIUser
	err := db.Model(User{}).Where("roleId=?", u.RoleId).Scan(&result).Error
	return result, err
}

func (u User) UpdateRole(db *gorm.DB) error {
	return db.Model(&u).Update("roleId", u.RoleId).Error
}
