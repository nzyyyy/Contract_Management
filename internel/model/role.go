package model

import "gorm.io/gorm"

type Role struct {
	ID          int `gorm:"autoIncrement"`
	Name        string
	Description string
	Functions   string
}

func (r Role) TableName() string {
	return "role"
}

func (r Role) Create(db *gorm.DB) error {
	return db.Create(&r).Error
}
func (r Role) Delete(db *gorm.DB) error {
	return db.Delete(&r).Error
}
