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
type RoleWithFunctionsResponse struct {
	RoleName    string   `json:"role_name"`
	Description string   `json:"description"`
	Functions   []string `json:"functions"`
}

func (svc *Service) CreateRole(params *CreateRoleRequest) error {
	return svc.dao.CreateRole(params.Name, params.Description, params.Functions)
}

func (svc *Service) DeleteRole(param int) error {
	return svc.dao.DeleteRole(param)
}
func (svc *Service) GetAllRole() ([]*model.APIRole, error) {
	return svc.dao.GetAllRole()
}
func (svc *Service) GetRoleById(id int) (*model.Role, error) {
	return svc.dao.GetRoleById(id)
}

func (svc *Service) GetRoleWithFunctionsById(id int) (*RoleWithFunctionsResponse, error) {
	role, err := svc.dao.GetRoleById(id)
	if err != nil {
		return nil, err
	}
	functions, err := svc.dao.GetFunctionsByRoleId(id)
	if err != nil {
		return nil, err
	}
	result := &RoleWithFunctionsResponse{}
	result.RoleName = role.Name
	result.Description = role.Description
	for _, function := range functions {
		result.Functions = append(result.Functions, function.Name)
	}
	return result, nil
}
func (svc *Service) GetFunctionsById(id int) ([]int, error) {
	//role, err := svc.dao.GetRoleById(id)
	//if err != nil {
	//	return nil, err
	//}
	functions, err := svc.dao.GetFuncByRoleId(id)
	if err != nil {
		return nil, err
	}
	var result []int
	for _, function := range functions {
		result = append(result, function.FunctionId)
	}
	return result, nil
}
