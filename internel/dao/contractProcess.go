package dao

import (
	"contract_management/internel/model"
	"fmt"
)

func (d Dao) DeleteContractProcess(id int) error {
	con := model.ContractProcess{ID: id}
	return con.Delete(d.engine)
}

func (d Dao) UpdateContractProcess(ID, contractId, Type, state int, content string) error {
	con := model.ContractProcess{ID: ID, ContractId: contractId, Type: Type, State: state, Content: content}
	err := con.Update(d.engine)
	if err != nil {
		return err
	}
	count, err := con.Count(d.engine)
	if err != nil {
		return err
	}
	if count == 0 {
		switch con.Type {
		case 1:
			c := model.ContractState{ContractId: con.ContractId, Type: 2}
			return c.Update(d.engine)
		case 2:
			c := model.ContractState{ContractId: con.ContractId, Type: 4}
			return c.Update(d.engine)
		case 3:
			c := model.ContractState{ContractId: con.ContractId, Type: 5}
			return c.Update(d.engine)
		default:
			return fmt.Errorf("参数错误")
		}
	}
	return nil
}

func (d Dao) CreateContractProcess(ContractID, Type, UserId int) error {
	c := model.ContractProcess{ContractId: ContractID, Type: Type, State: 0, UserId: UserId}
	return c.Create(d.engine)
}
