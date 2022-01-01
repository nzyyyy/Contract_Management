package model

import "gorm.io/gorm"

type Right struct {
	RoleId     int `gorm:"column:roleId" json:"roleId"`
	FunctionId int `gorm:"column:functionId" json:"functionId"`
}

func (r Right) TableName() string {
	return "rights"
}

type FunctionName struct {
	Name string `json:"name" gorm:"column:name"`
}

func (r Right) Create(db *gorm.DB) error {
	return db.Create(&r).Error
}
func (r Right) Delete(db *gorm.DB) error {
	return db.Delete(&r).Error
}
func (r Right) ListById(db *gorm.DB) ([]*FunctionName, error) {
	var result []*FunctionName
	err := db.Model(Right{}).Select("function.name").Where("roleId=?", r.RoleId).
		Joins("join function on rights.functionId=function.id").Scan(&result).Error
	return result, err
}
func (r Right) ListFuncIdById(db *gorm.DB) ([]*Right, error) {
	var result []*Right
	err := db.Model(Right{}).Select("functionId").Where("roleId=?", r.RoleId).Find(&result).Error
	return result, err
}

func (r Right) RoleIdExistFunc(db *gorm.DB) ([]*Right, error) {
	var result []*Right
	err := db.Select("roleId").Where("functionId=?", r.FunctionId).Find(&result).Error
	return result, err
}
