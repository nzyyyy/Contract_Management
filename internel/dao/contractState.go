package dao

import "contract_management/internel/model"

func (d Dao) UpdateContractState(Id, Type int) error {
	con := model.ContractState{ContractId: Id, Type: Type}
	return con.Update(d.engine)
}
