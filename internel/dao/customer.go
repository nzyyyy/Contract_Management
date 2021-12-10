package dao

import "contract_management/internel/model"

func (d Dao) CreateCustomer(Name, Address, Tel, Fax, Code, Bank, Account string) error {
	customer := model.Customer{Name: Name, Address: Address, Tel: Tel, Fax: Fax, Code: Code, Bank: Bank, Account: Account}
	return customer.Create(d.engine)
}

func (d Dao) DeleteCustomer(ID int) error {
	customer := model.Customer{ID: ID}
	return customer.Delete(d.engine)
}
