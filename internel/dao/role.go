package dao

import "contract_management/internel/model"

func (d Dao) CreateRole(name, description string, functions []*int) error {
	role := model.Role{Name: name, Description: description}
	id, err := role.Create(d.engine)
	if err != nil {
		return err
	}
	tx := d.engine.Begin()
	for _, function := range functions {
		right := model.Right{
			RoleId:     id,
			FunctionId: *function,
		}
		err := right.Create(tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func (d Dao) DeleteRole(id int) error {
	role := model.Role{ID: id}
	return role.Delete(d.engine)
}

func (d Dao) GetAllRole() ([]*model.APIRole, error) {
	return model.Role{}.List(d.engine)
}

func (d Dao) GetRoleById(id int) (*model.Role, error) {
	role := model.Role{ID: id}
	return role.RoleById(d.engine)
}
