package model

import (
	"gorm.io/gorm"
	"time"
)

type ContractState struct {
	ContractId int       `gorm:"column:con_id;primaryKey" json:"contractId"`
	Type       int       `json:"type"`
	Time       time.Time `gorm:"autoUpdateTime" json:"time"`
}

func (ContractState) TableName() string {
	return "contract_state"
}

func (c ContractState) Create(db *gorm.DB) error {
	return db.Create(&c).Error
}
func (c ContractState) Delete(db *gorm.DB) error {
	return db.Delete(&c).Error
}
func (c ContractState) Update(db *gorm.DB) error {
	return db.Model(&c).Update("type", c.Type).Error
}
