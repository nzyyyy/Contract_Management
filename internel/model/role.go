package model

import "gorm.io/gorm"

type Role struct {
	ID          int    `gorm:"autoIncrement" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type APIRole struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (r Role) TableName() string {
	return "role"
}

func (r Role) Create(db *gorm.DB) (int, error) {
	err := db.Create(&r).Error
	return r.ID, err
}
func (r Role) Delete(db *gorm.DB) error {
	return db.Delete(&r).Error
}
func (r Role) List(db *gorm.DB) ([]*APIRole, error) {
	var result []*APIRole
	err := db.Model(Role{}).Find(&result).Error
	return result, err
}
func (r Role) RoleById(db *gorm.DB) (*Role, error) {
	role := new(Role)
	err := db.First(&role, r.ID).Error
	return role, err
}
