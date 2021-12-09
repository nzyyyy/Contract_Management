package model

import "time"

type Log struct {
	ID      int       `gorm:"autoIncrement"`
	UserId  int       `gorm:"column:user_id"`
	Time    time.Time `gorm:"autoCreateTime"`
	Context string
}

func (l Log) TableName() string {
	return "log"
}
