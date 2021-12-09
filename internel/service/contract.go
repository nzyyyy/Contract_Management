package service

import (
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
	ID int `json:"id"`
}
type UpdateContractRequest struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
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
	return svc.dao.UpdateContract(params.ID, params.Content)
}
