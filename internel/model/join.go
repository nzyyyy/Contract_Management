package model

import "gorm.io/gorm"

type ContractWithState struct {
	ID   int
	Name string
	Type int
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
	ID       int
	Username string
	Email    string
	Name     string
	RoleId   int `gorm:"column:role_id"`
}

func (u UserWithRole) List(db *gorm.DB) ([]*UserWithRole, error) {
	var result []*UserWithRole
	err := db.Model(User{}).Select("user.id,user.username,user.email,role.name,rights.role_id").
		Joins("left join (rights join role on rights.role_id = role.id) on user.id = rights.user_id").
		Scan(&result).Error
	return result, err
}
