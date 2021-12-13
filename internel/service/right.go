package service

type CreateOrUpdateRightRequest struct {
	UserId      int    `json:"userId"`
	RoleId      int    `json:"roleId"`
	Description string `json:"description"`
	OperatorId  int    `json:"operatorId"`
}

type DeleteRightRequest struct {
	UserId int `json:"userId"`
	OperatorId int `json:"operatorId"`
}

func (svc *Service) CreateRight(params *CreateOrUpdateRightRequest) error {
	return svc.dao.CreateRight(params.UserId, params.RoleId, params.Description)
}

func (svc *Service) DeleteRight(param int) error {
	return svc.dao.DeleterRight(param)
}

func (svc *Service) UpdateRight(params *CreateOrUpdateRightRequest) error {
	return svc.dao.UpdateRight(params.UserId, params.RoleId, params.Description)
}
