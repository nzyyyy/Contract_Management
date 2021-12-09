package service

type DeleteContractProcessRequest struct {
	ID int
}

type UpdateContractProcessRequest struct {
	ID         int    `json:"id"`
	ContractId int    `json:"contractId"`
	Type       int    `json:"type"`
	State      int    `json:"state"`
	Content    string `json:"content"`
}

func (svc *Service) DeleteContractProcess(params *DeleteContractProcessRequest) error {
	return svc.dao.DeleteContractProcess(params.ID)
}

func (svc *Service) UpdateContractProcess(params *UpdateContractProcessRequest) error {
	return svc.dao.UpdateContractProcess(params.ID, params.ContractId, params.Type, params.State, params.Content)
}
