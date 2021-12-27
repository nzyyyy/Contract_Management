package dao

import (
	"contract_management/internel/model"
	"time"
)

func (d Dao) CreateContract(name string, customerId int, beginTime, endTime time.Time, content string, userId int, filename, path, Type string) error {
	contract := model.Contract{
		Name:       name,
		CustomerId: customerId,
		BeginTime:  beginTime,
		EndTime:    endTime,
		Content:    content,
		UserId:     userId,
	}
	id, err := contract.Create(d.engine)
	if err != nil {
		return err
	}
	contractAttachment := model.ContractAttachment{ContractId: id, FileName: filename, Path: path, Type: Type}
	return contractAttachment.Create(d.engine)
}
func (d Dao) DeleteContract(ID int) error {
	contract := model.Contract{ID: ID}
	return contract.Delete(d.engine)
}

func (d Dao) UpdateContract(id int, content string) error {
	contract := model.Contract{ID: id, Content: content}
	return contract.Update(d.engine)
}
func (d Dao) GetAllList() ([]*model.APIContract, error) {
	return model.Contract{}.AllList(d.engine)
}
func (d Dao) GetListByUserID(userId int) ([]*model.APIContract, error) {
	return model.Contract{UserId: userId}.ListByUser(d.engine)
}
func (d Dao) GetContractById(id int) (*model.Contract, error) {
	c := model.Contract{ID: id}
	return c.GetContractById(d.engine)
}
