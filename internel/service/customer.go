package service

import "contract_management/internel/model"

type CreateCustomerRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Tel     string `json:"tel"`
	Fax     string `json:"fax"`
	Code    string `json:"code"`
	Bank    string `json:"bank"`
	Account string `json:"account"`
	UserId  int    `json:"userId"`
}
type DeleteCustomerRequest struct {
	CustomerId int `json:"customerId"`
	UserId     int `json:"userId"`
}

func (svc *Service) CreateCustomer(params *CreateCustomerRequest) error {
	return svc.dao.CreateCustomer(params.Name, params.Address, params.Tel, params.Fax, params.Code, params.Bank, params.Account)
}

func (svc *Service) DeleteCustomer(param int) error {
	return svc.dao.DeleteCustomer(param)
}
func (svc *Service) GetCustomerPartList() ([]*model.APICustomer, error) {
	return svc.dao.GetCustomerPartList()
}
func (svc *Service) GetCustomerById(ID int) (*model.Customer, error) {
	return svc.dao.GetOneCustomerById(ID)
}
