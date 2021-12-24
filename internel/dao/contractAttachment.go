package dao

import "contract_management/internel/model"

func (d Dao) UpdateFilePath(contractId int, path string) error {
	c := model.ContractAttachment{ContractId: contractId, Path: path}
	return c.Update(d.engine)
}
