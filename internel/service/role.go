package service

type CreateRoleRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Functions   string `json:"functions"`
	OperatorId 	int `json:"operatorId"`
}

type DeleteRoleRequest struct {
	RoleId int `json:"roleId"`
	OperatorId int `json:"operatorId"`
}

func (svc *Service) CreateRole(params *CreateRoleRequest) error {
	return svc.dao.CreateRole(params.Name, params.Description, params.Functions)
}

func (svc *Service) DeleteRole(param int) error {
	return svc.dao.DeleteCustomer(param)
}
