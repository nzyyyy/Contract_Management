package service

import "contract_management/internel/model"

type CreateRoleRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Functions   []*int `json:"functions"`
	OperatorId  int    `json:"operatorId"`
}

type DeleteRoleRequest struct {
	RoleId     int `json:"roleId"`
	OperatorId int `json:"operatorId"`
}

func (svc *Service) CreateRole(params *CreateRoleRequest) error {
	return svc.dao.CreateRole(params.Name, params.Description, params.Functions)
}

func (svc *Service) DeleteRole(param int) error {
	return svc.dao.DeleteCustomer(param)
}
func (svc *Service) GetAllRole() ([]*model.APIRole, error) {
	return svc.dao.GetAllRole()
}
func (svc *Service) GetRoleById(id int) (*model.Role, error) {
	return svc.dao.GetRoleById(id)
}
