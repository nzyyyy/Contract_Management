package model

import (
	"gorm.io/gorm"
	"time"
)

type ContractTimeLine struct {
	ContractId int       `gorm:"column:contractId" json:"contractId"`
	Content    string    `json:"content"`
	Time       time.Time `json:"time" gorm:"autoUpdateTime"`
}

func (c ContractTimeLine) TableName() string {
	return "contract_timeline"
}

func (c ContractTimeLine) Create(db *gorm.DB) error {
	return db.Create(&c).Error
}

func (c ContractTimeLine) GetById(db *gorm.DB) ([]*ContractTimeLine, error) {
	var result []*ContractTimeLine
	err := db.Model(ContractTimeLine{}).Where("contractId=?", c.ContractId).Find(&result).Error
	return result, err
}
