package model

import "gorm.io/gorm"

type Customer struct {
	ID      int `gorm:"autoIncrement"`
	Name    string
	Address string
	Tel     string
	Fax     string
	Code    string
	Bank    string
	Account string
}

func (Customer) TableName() string {
	return "customer"
}

func (c Customer) Create(db *gorm.DB) error {
	return db.Create(&c).Error
}

func (c Customer) Delete(db *gorm.DB) error {
	return db.Delete(&c).Error
}
