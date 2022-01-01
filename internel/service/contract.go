package service

import (
	"contract_management/internel/model"
	"contract_management/pkg/util"
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
	id, err := svc.dao.CreateContract(params.Name, params.CustomerID,
		beginTime, endTime, params.Content, params.UserId,
		params.FileName, params.Path, params.Type)
	if err != nil {
		return err
	}
	err = svc.dao.CreateTimeLine(id, "起草合同完成")
	if err != nil {
		return err
	}
	return nil
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
	go func() {
		operators, err := svc.dao.SelectOperator(params.ID, 2)
		if err != nil || len(operators) == 0 {
			return
		}
		var emails []Email
		for _, operator := range operators {
			emails = append(emails, Email{ID: operator, Content: "您有新的合同在等待审批"})
		}
		util.PostRequest("http://localhost:8001/email", map[string]interface{}{"emails": emails})
	}()
	err = svc.dao.CreateTimeLine(params.ID, "定稿合同完成")
	if err != nil {
		return err
	}
	return nil
}
func (svc *Service) GetAllList() ([]*model.APIContract, error) {
	return svc.dao.GetAllList()
}
func (svc *Service) GetListByUser(userId int) ([]*model.APIContract, error) {
	return svc.dao.GetListByUserID(userId)
}
func (svc *Service) GetContractById(id int) (*model.ContractWithFull, error) {
	return svc.dao.GetContractById(id)
}
func (svc *Service) GetContractWithoutOperator() ([]*model.APIContract, error) {
	return svc.dao.GetContractWithoutOperator()
}
