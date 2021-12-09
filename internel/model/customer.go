package model

type Customer struct {
	ID      int `gorm:"autoIncrement"`
	Name    string
	Address string
	Tel     string
	Fax     string
	Code    string
	Bank    string
	Account string
}

func (Customer) TableName() string {
	return "customer"
}
