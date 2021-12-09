package model

import (
	"time"

	"gorm.io/gorm"
)

type Contract struct {
	ID         int `gorm:"autoIncrement"`
	Name       string
	CustomerId int       `gorm:"column:customer_id"`
	BeginTime  time.Time `gorm:"column:BeginTime"`
	EndTime    time.Time `gorm:"column:EndTime"`
	Content    string
	UserId     int `gorm:"column:user_id"`
}

func (Contract) TableName() string {
	return "contract"
}

func (c Contract) Create(db *gorm.DB) (int, error) {
	err := db.Create(&c).Error
	if err != nil {
		return 0, err
	}
	return c.ID, nil
}
func (c Contract) Delete(db *gorm.DB) error {
	return db.Delete(&c).Error
}
func (c Contract) Update(db *gorm.DB) error {
	return db.Model(&c).Update("content", c.Content).Error
}

func (c Contract) AfterCreate(db *gorm.DB) error {
	con := ContractState{
		ContractId: c.ID,
		Type:       1,
	}
	return con.Create(db)
}
