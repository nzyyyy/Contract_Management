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
	UserId     int `json:"userId"`
}
type CreateContractProcessRequest struct {
	ContractId int `json:"contractId"`
	Processes []Process `json:"process"`
	OperatorId int `json:"operatorId"`
}

type Process struct {
	Type       int `json:"type"`
	UserId     int `json:"userId"`
}

func (svc *Service) DeleteContractProcess(params *DeleteContractProcessRequest) error {
	return svc.dao.DeleteContractProcess(params.ID)
}

func (svc *Service) UpdateContractProcess(params *UpdateContractProcessRequest) error {
	return svc.dao.UpdateContractProcess(params.ID, params.ContractId, params.Type, params.State, params.Content)
}

func (svc *Service) CreateContractProcess(params *CreateContractProcessRequest) error {
	for _,process := range params.Processes{
		if err := svc.dao.CreateContractProcess(params.ContractId, process.Type, process.UserId); err != nil {
			return err
		}
	}
	return nil
}
