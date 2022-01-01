package model

import (
	"time"

	"gorm.io/gorm"
)

type Contract struct {
	ID         int       `gorm:"autoIncrement" json:"id"`
	Name       string    `json:"name"`
	CustomerId int       `gorm:"column:customer_id" json:"customerId"`
	BeginTime  time.Time `gorm:"column:begin_time" json:"beginTime"`
	EndTime    time.Time `gorm:"column:end_time" json:"endTime"`
	Content    string    `json:"content"`
	UserId     int       `gorm:"column:user_id" json:"userId"`
}

func (Contract) TableName() string {
	return "contract"
}

type APIContract struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type ContractWithFull struct {
	ID           int       `gorm:"autoIncrement" json:"id"`
	Name         string    `json:"name" gorm:"column:name"`
	CustomerId   int       `gorm:"column:customer_id" json:"customerId"`
	CustomerName string    `gorm:"column:customer_name" json:"customerName"`
	BeginTime    time.Time `gorm:"column:begin_time" json:"beginTime"`
	EndTime      time.Time `gorm:"column:end_time" json:"endTime"`
	Content      string    `json:"content"`
	UserId       int       `gorm:"column:user_id" json:"userId"`
	Username     string    `json:"username" gorm:"column:username"`
	FilePath     string    `json:"filePath" gorm:"column:path"`
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
func (c Contract) AllList(db *gorm.DB) ([]*APIContract, error) {
	var result []*APIContract
	err := db.Model(Contract{}).Find(&result).Error
	return result, err
}
func (c Contract) ListByUser(db *gorm.DB) ([]*APIContract, error) {
	var result []*APIContract
	err := db.Model(Contract{}).Where("user_id = ?", c.UserId).Find(&result).Error
	return result, err
}
func (c Contract) ListWithoutOperator(db *gorm.DB) ([]*APIContract, error) {
	var result []*APIContract
	err := db.Model(Contract{}).Where("id not in (?)",
		db.Table("contract_process").Select("con_id")).Scan(&result).Error
	return result, err
}
func (c Contract) GetContractById(db *gorm.DB) (*ContractWithFull, error) {
	var result = new(ContractWithFull)
	err := db.Model(Contract{}).Select("contract.id,contract.name,contract.customer_id,customer.name as customer_name,contract.begin_time,contract.end_time,contract.content,contract.user_id,user.username,contract_attachment.path").
		Joins("join customer on contract.customer_id=customer.id").
		Joins("join user on contract.user_id=user.id").
		Joins("join contract_attachment on contract.id=contract_attachment.con_id").
		Where("contract.id=?", c.ID).
		Scan(&result).Error
	return result, err
}
func (c Contract) AfterCreate(db *gorm.DB) error {
	con := ContractState{
		ContractId: c.ID,
		Type:       1,
	}
	return con.Create(db)
}
