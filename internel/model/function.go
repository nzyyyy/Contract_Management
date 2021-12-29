package model

import "gorm.io/gorm"

type Function struct {
	ID   int    `gorm:"autoIncrement" json:"id"`
	Name string `json:"name"`
}

func (Function) TableName() string {
	return "function"
}

func (f Function) List(db *gorm.DB) ([]*Function, error) {
	var result []*Function
	err := db.Find(&result).Error
	return result, err
}
