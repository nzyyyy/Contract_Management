package service

import (
	"contract_management/internel/model"
	"contract_management/pkg/util"
)

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
	err := svc.dao.UpdateContractProcess(params.ID, params.State, params.Content)
	if err != nil {
		return err
	}
	count := svc.dao.CountProcess(params.ContractId, params.Type, 0)
	if count == 0 {
		switch params.Type {
		case 1:
			err = svc.dao.UpdateContractState(params.ContractId, 2)
			if err != nil {
				return err
			}
			err = svc.dao.CreateTimeLine(params.ContractId, "合同会签完成")
			if err != nil {
				return err
			}
			go func() {
				con, err2 := svc.dao.GetContractById(params.ContractId)
				if err2 != nil {
					return
				}
				var emails []Email
				emails = append(emails, Email{ID: con.UserId, Content: "您有合同已完成会签,请及时定稿"})
				util.PostRequest("http://localhost:8001/email", map[string]interface{}{"emails": emails})
			}()
		case 2:
			temp := svc.dao.CountProcess(params.ContractId, 2, 2)
			if temp != 0 {
				err = svc.dao.UpdateContractState(params.ContractId, 6)
				if err != nil {
					return err
				}
				err = svc.dao.CreateTimeLine(params.ContractId, "合同审批不通过,合同终止")
				if err != nil {
					return err
				}
			} else {
				err = svc.dao.UpdateContractState(params.ContractId, 4)
				if err != nil {
					return err
				}
				err = svc.dao.CreateTimeLine(params.ContractId, "合同审批完成")
				if err != nil {
					return err
				}
				go func() {
					operators, err := svc.dao.SelectOperator(params.ContractId, 3)
					if err != nil || len(operators) == 0 {
						return
					}
					var emails []Email
					for _, operator := range operators {
						emails = append(emails, Email{ID: operator, Content: "您有新的合同在等待签订"})
					}
					util.PostRequest("http://localhost:8001/email", map[string]interface{}{"emails": emails})
				}()
			}
		case 3:
			err = svc.dao.UpdateContractState(params.ContractId, 5)
			if err != nil {
				return err
			}
			err = svc.dao.CreateTimeLine(params.ContractId, "合同签订完成")
			if err != nil {
				return err
			}
			go func() {
				con, err2 := svc.dao.GetContractById(params.ContractId)
				if err2 != nil {
					return
				}
				var emails []Email
				emails = append(emails, Email{ID: con.UserId, Content: "您的合同已完成签订"})
				util.PostRequest("http://localhost:8001/email", map[string]interface{}{"emails": emails})
			}()
		}
	}
	return nil
}

func (svc *Service) CreateContractProcess(params *CreateContractProcessRequest) error {
	tx := svc.dao.GetTX()
	for _, process := range params.Processes {
		if err := tx.CreateContractProcess(params.ContractId, process.Type, process.UserId); err != nil {
			tx.RollBack()
			return err
		}
	}
	tx.Commit()
	err := svc.dao.CreateTimeLine(params.ContractId, "合同正在处理")
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) SelectContractComment(ContractId, Type int) ([]*model.CommentOfContract, error) {
	return svc.dao.SelectContractComment(ContractId, Type)
}
