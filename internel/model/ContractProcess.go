package model

import (
	"time"

	"gorm.io/gorm"
)

type ContractProcess struct {
	ID         int       `gorm:"autoIncrement" json:"id"`
	ContractId int       `gorm:"column:con_id" json:"contractId"`
	Type       int       `json:"type"`
	State      int       `json:"state"`
	UserId     int       `gorm:"column:user_id" json:"userId"`
	Content    string    `json:"content"`
	Time       time.Time `gorm:"autoUpdateTime" json:"time"`
}
type CommentOfContract struct {
	Content string    `json:"content"`
	Time    time.Time `json:"time"`
}

func (ContractProcess) TableName() string {
	return "contract_process"
}
func (c ContractProcess) Update(db *gorm.DB) error {
	return db.Model(&c).Updates(map[string]interface{}{"state": c.State, "content": c.Content}).Error
}
func (c ContractProcess) Delete(db *gorm.DB) error {
	return db.Delete(&c).Error
}
func (c ContractProcess) Count(db *gorm.DB) (int64, error) {
	var count int64
	db.Model(&ContractProcess{}).Where("con_id=? AND type=? AND state=?", c.ContractId, c.Type, c.State).Count(&count)
	return count, nil
}
func (c ContractProcess) Create(db *gorm.DB) error {
	return db.Create(&c).Error
}
func (c ContractProcess) SelectContractComment(db *gorm.DB) ([]*CommentOfContract, error) {
	var result []*CommentOfContract
	err := db.Model(ContractProcess{}).Where("con_id=? and type=?", c.ContractId, c.Type).Find(&result).Error
	return result, err
}
