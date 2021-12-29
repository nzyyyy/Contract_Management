package dao

import "contract_management/internel/model"

func (d Dao) CreateTimeLine(conId int, content string) error {
	line := model.ContractTimeLine{ContractId: conId, Content: content}
	return line.Create(d.engine)
}

func (d Dao) GetTimeLineByID(conId int) ([]*model.ContractTimeLine, error) {
	line := model.ContractTimeLine{ContractId: conId}
	return line.GetById(d.engine)
}
