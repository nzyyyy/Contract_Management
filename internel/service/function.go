package service

import "contract_management/internel/model"

func (svc *Service) GetFunctionList() ([]*model.Function, error) {
	return svc.dao.GetFunctionList()
}
