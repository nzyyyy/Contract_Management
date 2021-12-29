package service

import "contract_management/internel/model"

func (svc *Service) GetTimeLineById(id int) ([]*model.ContractTimeLine, error) {
	return svc.dao.GetTimeLineByID(id)
}
