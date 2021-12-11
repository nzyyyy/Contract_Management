package dao

import "contract_management/internel/model"

func (d Dao) CreateLog(userId int, content string) error {
	log := model.Log{UserId: userId, Context: content}
	return log.Create(d.engine)
}

func (d Dao) DeleteLog(id int) error {
	log := model.Log{ID: id}
	return log.Delete(d.engine)
}
