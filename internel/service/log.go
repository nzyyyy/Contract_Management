package service

import "contract_management/internel/model"

type CreateLogRequest struct {
	UserId  int
	Content string
}
type DeleteLogRequest struct {
	UserId int
	LogId  int
}

func (svc *Service) CreateLog(params *CreateLogRequest) error {
	return svc.dao.CreateLog(params.UserId, params.Content)
}

func (svc *Service) DeleteLog(id int) error {
	return svc.dao.DeleteLog(id)
}
func (svc *Service) GetLogList() ([]*model.APILog, error) {
	return svc.dao.GetLogList()
}
