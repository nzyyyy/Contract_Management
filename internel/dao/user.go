package dao

import "contract_management/internel/model"

func (d Dao) GetUserByRole(roleId int) ([]*model.APIUser, error) {
	return model.User{RoleId: roleId}.ListByRole(d.engine)
}

func (d Dao) RemoveUserRole(userId, roleId int) error {
	return model.User{ID: userId, RoleId: roleId}.UpdateRole(d.engine)
}
