package service

import "contract_management/internel/model"

func (svc *Service) GetAllProcessList() ([]*model.ContractWithState, error) {
	return svc.dao.GetAllProcessList()
}

func (svc *Service) GetProcessListByUser(id int) ([]*model.ContractWithState, error) {
	return svc.dao.GetAllProcessListByUser(id)
}
func (svc *Service) GetUserWithRole() ([]*model.UserWithRole, error) {
	return svc.dao.GetAllUserWithRole()
}
