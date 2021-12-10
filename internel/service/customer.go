package service

type CreateCustomerRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Tel     string `json:"tel"`
	Fax     string `json:"fax"`
	Code    string `json:"code"`
	Bank    string `json:"bank"`
	Account string `json:"account"`
}

func (svc *Service) CreateCustomer(params *CreateCustomerRequest) error {
	return svc.dao.CreateCustomer(params.Name, params.Address, params.Tel, params.Fax, params.Code, params.Bank, params.Account)
}

func (svc *Service) DeleteCustomer(param int) error {
	return svc.dao.DeleteCustomer(param)
}
