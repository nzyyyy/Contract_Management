package dao

import "contract_management/internel/model"

func (d Dao) GetFunctionList() ([]*model.Function, error) {
	return model.Function{}.List(d.engine)
}
