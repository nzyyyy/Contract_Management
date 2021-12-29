package model

import "gorm.io/gorm"

type Right struct {
	RoleId     int `gorm:"column:roleId" json:"roleId"`
	FunctionId int `gorm:"column:functionId" json:"functionId"`
}

func (r Right) TableName() string {
	return "rights"
}

func (r Right) Create(db *gorm.DB) error {
	return db.Create(&r).Error
}
func (r Right) Delete(db *gorm.DB) error {
	return db.Delete(&r).Error
}
