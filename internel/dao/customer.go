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

func (d Dao) GetCustomerPartList() ([]*model.APICustomer, error) {
	customer := model.Customer{}
	return customer.APIList(d.engine)
}
func (d Dao) GetOneCustomerById(ID int) (*model.Customer, error) {
	customer := model.Customer{ID: ID}
	c, err := customer.GetOneCustomer(d.engine)
	return c, err
}
