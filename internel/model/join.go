package model

import "gorm.io/gorm"

type ContractWithState struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type int    `json:"type"`
}

func (c ContractWithState) AllList(db *gorm.DB) ([]*ContractWithState, error) {
	var result []*ContractWithState
	err := db.Model(Contract{}).Select("contract.id,contract.name,contract_state.type").
		Joins("join contract_state on contract.id=contract_state.con_id").
		Scan(&result).Error
	return result, err
}

func (c ContractWithState) ListByUser(db *gorm.DB, id int) ([]*ContractWithState, error) {
	var result []*ContractWithState
	err := db.Model(Contract{}).Select("contract.id,contract.name,contract_state.type").
		Joins("join contract_state on contract.id=contract_state.con_id").
		Where("user_id=?", id).
		Scan(&result).Error
	return result, err
}

type UserWithRole struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

func (u UserWithRole) List(db *gorm.DB) ([]*UserWithRole, error) {
	var result []*UserWithRole
	err := db.Model(User{}).Select("user.id,user.username,user.email,role.name").
		Joins("left join role on user.roleId=role.id").
		Scan(&result).Error
	return result, err
}
