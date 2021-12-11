package dao

import "contract_management/internel/model"

func (d Dao) CreateRole(name, description, functions string) error {
	role := model.Role{Name: name, Description: description, Functions: functions}
	return role.Create(d.engine)
}

func (d Dao) DeleteRole(id int) error {
	role := model.Role{ID: id}
	return role.Delete(d.engine)
}
