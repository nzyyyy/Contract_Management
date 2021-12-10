package dao

import "contract_management/internel/model"

func (d Dao) CreateRight(userId, roleId int, description string) error {
	r := model.Right{
		UserId:      userId,
		RoleId:      roleId,
		Description: description,
	}
	return r.Create(d.engine)
}

func (d Dao) DeleterRight(userId int) error {
	r := model.Right{UserId: userId}
	return r.Delete(d.engine)
}

func (d Dao) UpdateRight(userId, roleId int, description string) error {
	r := model.Right{
		UserId:      userId,
		RoleId:      roleId,
		Description: description,
	}
	return r.Update(d.engine)
}
