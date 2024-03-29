package model

import "gorm.io/gorm"

type Customer struct {
	ID      int    `gorm:"autoIncrement" json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Tel     string `json:"tel"`
	Fax     string `json:"fax"`
	Code    string `json:"code"`
	Bank    string `json:"bank"`
	Account string `json:"account"`
}

type APICustomer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (Customer) TableName() string {
	return "customer"
}

func (c Customer) Create(db *gorm.DB) error {
	return db.Create(&c).Error
}

func (c Customer) Delete(db *gorm.DB) error {
	return db.Delete(&c).Error
}

func (c Customer) APIList(db *gorm.DB) ([]*APICustomer, error) {
	var result []*APICustomer
	err := db.Model(Customer{}).Find(&result).Error
	return result, err
}

func (c Customer) GetOneCustomer(db *gorm.DB) (*Customer, error) {
	customer := new(Customer)
	err := db.First(&customer, c.ID).Error
	return customer, err
}
func (c Customer) Update(db *gorm.DB) error {
	return db.Save(&c).Error
}
