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
	UserId     int    `json:"userId"`
}
type CreateContractProcessRequest struct {
	ContractId int       `json:"contractId"`
	Processes  []Process `json:"process"`
	OperatorId int       `json:"operatorId"`
}

type Process struct {
	Type   int `json:"type"`
	UserId int `json:"userId"`
}

func (svc *Service) DeleteContractProcess(params *DeleteContractProcessRequest) error {
	return svc.dao.DeleteContractProcess(params.ID)
}

func (svc *Service) UpdateContractProcess(params *UpdateContractProcessRequest) error {
	//TODO 会签，审批，签订
	err := svc.dao.UpdateContractProcess(params.ID, params.State, params.Content)
	if err != nil {
		return err
	}
	count := svc.dao.CountProcess(params.ContractId, params.Type, 0)
	if count == 0 {
		switch params.Type {
		case 1:
			_ = svc.dao.UpdateContractState(params.ContractId, 2)
			//TODO 通知起草人定稿合同
		case 2:
			temp := svc.dao.CountProcess(params.ContractId, 2, 2)
			if temp != 0 {
				_ = svc.dao.UpdateContractState(params.ContractId, 6)
			} else {
				_ = svc.dao.UpdateContractState(params.ContractId, 4)
				//TODO 通知签订人签订合同
			}
		case 3:
			_ = svc.dao.UpdateContractState(params.ContractId, 5)
			//TODO 通知起草人合同签订完成
		}
	}
	return nil
}

func (svc *Service) CreateContractProcess(params *CreateContractProcessRequest) error {
	for _, process := range params.Processes {
		if err := svc.dao.CreateContractProcess(params.ContractId, process.Type, process.UserId); err != nil {
			return err
		}
	}
	return nil
}
