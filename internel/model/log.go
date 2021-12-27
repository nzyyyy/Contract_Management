package model

import (
	"gorm.io/gorm"
	"time"
)

type Log struct {
	ID      int       `gorm:"autoIncrement" json:"id"`
	UserId  int       `gorm:"column:user_id" json:"userId"`
	Time    time.Time `gorm:"autoCreateTime" json:"time"`
	Context string    `json:"context"`
}

func (l Log) TableName() string {
	return "log"
}

func (l Log) Create(db *gorm.DB) error {
	return db.Create(&l).Error
}

func (l Log) Delete(db *gorm.DB) error {
	return db.Delete(&l).Error
}
