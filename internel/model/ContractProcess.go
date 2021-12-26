package model

import (
	"time"

	"gorm.io/gorm"
)

type ContractProcess struct {
	ID         int `gorm:"autoIncrement"`
	ContractId int `gorm:"column:con_id"`
	Type       int
	State      int
	UserId     int `gorm:"column:user_id"`
	Content    string
	Time       time.Time `gorm:"autoUpdateTime"`
}
type CommentOfContract struct {
	Content string
	Time    time.Time
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
