package service

import (
	"contract_management/internel/model"
	"github.com/araddon/dateparse"
)

type CreateContractRequest struct {
	Name       string `json:"name"`
	CustomerID int    `json:"customerId"`
	BeginTime  string `json:"beginTime"`
	EndTime    string `json:"endTime"`
	Content    string `json:"content"`
	UserId     int    `json:"userId"`
	FileName   string `json:"fileName"`
	Path       string `json:"path"`
	Type       string `json:"type"`
}
type DeleteContractRequest struct {
	ID     int `json:"id"`
	UserId int `json:"userId"`
}
type UpdateContractRequest struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Path    string `json:"path"`
	UserId  int    `json:"userId"`
}

func (svc *Service) CreateContract(params *CreateContractRequest) error {
	beginTime, err := dateparse.ParseLocal(params.BeginTime)
	if err != nil {
		return err
	}
	endTime, err := dateparse.ParseLocal(params.EndTime)
	if err != nil {
		return err
	}
	return svc.dao.CreateContract(params.Name, params.CustomerID,
		beginTime, endTime, params.Content, params.UserId,
		params.FileName, params.Path, params.Type)
}

func (svc *Service) DeleteContract(params *DeleteContractRequest) error {
	return svc.dao.DeleteContract(params.ID)
}

func (svc *Service) UpdateContract(params *UpdateContractRequest) error {
	err := svc.dao.UpdateContract(params.ID, params.Content)
	if err != nil {
		return err
	}
	err = svc.dao.UpdateFilePath(params.ID, params.Path)
	if err != nil {
		return err
	}
	err = svc.dao.UpdateContractState(params.ID, 3)
	if err != nil {
		return err
	}
	//TODO 通知审批人审批合同
	return nil
}
func (svc *Service) GetAllList() ([]*model.APIContract, error) {
	return svc.dao.GetAllList()
}
func (svc *Service) GetListByUser(userId int) ([]*model.APIContract, error) {
	return svc.dao.GetListByUserID(userId)
}
