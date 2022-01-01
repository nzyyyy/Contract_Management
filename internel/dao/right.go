package dao

import "contract_management/internel/model"

func (d Dao) CreateRight(functionId, roleId int) error {
	r := model.Right{
		RoleId:     roleId,
		FunctionId: functionId,
	}
	return r.Create(d.engine)
}
func (d Dao) GetFunctionsByRoleId(roleId int) ([]*model.FunctionName, error) {
	return model.Right{RoleId: roleId}.ListById(d.engine)
}

func (d Dao) GetRoleIdByFunc(f int) ([]*model.Right, error) {
	return model.Right{FunctionId: f}.RoleIdExistFunc(d.engine)
}
func (d Dao) GetFuncByRoleId(id int) ([]*model.Right, error) {
	return model.Right{RoleId: id}.ListFuncIdById(d.engine)
}
