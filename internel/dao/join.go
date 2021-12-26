package dao

import "contract_management/internel/model"

func (d Dao) GetAllProcessList() ([]*model.ContractWithState, error) {
	return model.ContractWithState{}.AllList(d.engine)
}

func (d Dao) GetAllProcessListByUser(id int) ([]*model.ContractWithState, error) {
	return model.ContractWithState{}.ListByUser(d.engine, id)
}

func (d Dao) GetAllUserWithRole() ([]*model.UserWithRole, error) {
	return model.UserWithRole{}.List(d.engine)
}
