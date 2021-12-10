package model

import "gorm.io/gorm"

type Right struct {
	UserId      int `gorm:"column:user_id;primaryKey"`
	RoleId      int `gorm:"column:role_id"`
	Description string
}

func (r Right) TableName() string {
	return "right"
}

func (r Right) Create(db *gorm.DB) error {
	return db.Create(&r).Error
}
func (r Right) Delete(db *gorm.DB) error {
	return db.Delete(&r).Error
}
func (r Right) Update(db *gorm.DB) error {
	return db.Model(&r).Updates(map[string]interface{}{"role_id": r.RoleId, "description": r.Description}).Error
}
