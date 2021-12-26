package dao

import (
	"contract_management/internel/model"
)

func (d Dao) DeleteContractProcess(id int) error {
	con := model.ContractProcess{ID: id}
	return con.Delete(d.engine)
}

func (d Dao) UpdateContractProcess(ID, state int, content string) error {
	con := model.ContractProcess{ID: ID, State: state, Content: content}
	return con.Update(d.engine)
}

func (d Dao) CreateContractProcess(ContractID, Type, UserId int) error {
	c := model.ContractProcess{ContractId: ContractID, Type: Type, State: 0, UserId: UserId}
	return c.Create(d.engine)
}

func (d Dao) CountProcess(ContractID, Type, State int) int64 {
	c := model.ContractProcess{ContractId: ContractID, Type: Type, State: State}
	count, _ := c.Count(d.engine)
	return count
}

func (d Dao) SelectContractComment(ContractId, Type int) ([]*model.CommentOfContract, error) {
	c := model.ContractProcess{ContractId: ContractId, Type: Type}
	return c.SelectContractComment(d.engine)

}
