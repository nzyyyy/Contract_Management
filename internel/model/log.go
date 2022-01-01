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
type APILog struct {
	ID       int       `gorm:"autoIncrement" json:"id"`
	Username string    `gorm:"column:username" json:"username"`
	Time     time.Time `gorm:"autoCreateTime" json:"time"`
	Context  string    `json:"context"`
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

func (l Log) List(db *gorm.DB) ([]*APILog, error) {
	var result []*APILog
	err := db.Model(Log{}).Select("log.id,user.username,log.time,log.context").Joins("join user on log.user_id=user.id").Scan(&result).Error
	return result, err
}
